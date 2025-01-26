package server

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"cloud.google.com/go/spanner"
	pb "github.com/parampatil/stealth-backend/proto"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HealthServer struct {
    pb.UnimplementedHealthServiceServer
    spannerClient *spanner.Client
}

func NewHealthServer(spannerClient *spanner.Client) *HealthServer {
    return &HealthServer{spannerClient: spannerClient}
}

func (s *HealthServer) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
    return &pb.HealthCheckResponse{
        Status: "SERVING",
    }, nil
}

func (s *HealthServer) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
    message := "Hello, " + req.Name + "!"
    return &pb.GreetingResponse{
        Message: message,
    }, nil
}

func (s *HealthServer) GetEarnings(ctx context.Context, req *pb.GetEarningsRequest) (*pb.GetEarningsResponse, error) {
    var earnings []*pb.Earning
    stmt := spanner.Statement{
        SQL: `SELECT earning_id, amount, date FROM Earnings WHERE uid = @uid ORDER BY date DESC`,
        Params: map[string]interface{}{
            "uid": req.Uid,
        },
    }
    iter := s.spannerClient.Single().Query(ctx, stmt)
    defer iter.Stop()

    for {
        row, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, status.Errorf(codes.Internal, "Failed to query earnings: %v", err)
        }

        var earning pb.Earning
        var date spanner.NullTime
        if err := row.Columns(&earning.EarningId, &earning.Amount, &date); err != nil {
            return nil, status.Errorf(codes.Internal, "Failed to parse earnings: %v", err)
        }
        earning.Date = timestamppb.New(date.Time)
        earnings = append(earnings, &earning)
    }

    return &pb.GetEarningsResponse{Earnings: earnings}, nil
}

func (s *HealthServer) CreateEarning(ctx context.Context, req *pb.CreateEarningRequest) (*pb.CreateEarningResponse, error) {
    mutation := spanner.Insert(
        "Earnings",
        []string{"uid", "earning_id", "amount", "date"},
        []interface{}{req.Uid, req.Earning.EarningId, req.Earning.Amount, req.Earning.Date.AsTime()},
    )

    _, err := s.spannerClient.Apply(ctx, []*spanner.Mutation{mutation})
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to insert earning: %v", err)
    }

    return &pb.CreateEarningResponse{Status: "SUCCESS"}, nil
}

func (s *HealthServer) GenerateDemoData(ctx context.Context, req *pb.GenerateDemoDataRequest) (*pb.GenerateDemoDataResponse, error) {
    var mutations []*spanner.Mutation
    uid := req.Uid
    startDate := time.Now().AddDate(-3, 0, 0) // 3 years ago
    for i := 0; i < 36; i++ { // 36 months
        earningID := fmt.Sprintf("%s-%d", uid, i)
        amount := rand.Float64() * 1000 // Random amount between 0 and 1000
        date := startDate.AddDate(0, i, 0)
        mutation := spanner.Insert(
            "Earnings",
            []string{"uid", "earning_id", "amount", "date"},
            []interface{}{uid, earningID, amount, date},
        )
        mutations = append(mutations, mutation)
    }

    _, err := s.spannerClient.Apply(ctx, mutations)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to generate demo data: %v", err)
    }

    return &pb.GenerateDemoDataResponse{Status: "SUCCESS"}, nil
}