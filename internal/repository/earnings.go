package repository

import (
	"context"

	"cloud.google.com/go/spanner"
	pb "github.com/parampatil/stealth-backend/proto"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EarningsRepository struct {
    client *spanner.Client
}

func NewEarningsRepository(client *spanner.Client) *EarningsRepository {
    return &EarningsRepository{client: client}
}

func (r *EarningsRepository) GetEarnings(ctx context.Context, req *pb.GetEarningsRequest) (*pb.GetEarningsResponse, error) {
    var earnings []*pb.Earning
    stmt := spanner.Statement{
        SQL: `SELECT earning_id, amount, date FROM Earnings WHERE uid = @uid ORDER BY date DESC`,
        Params: map[string]interface{}{
            "uid": req.Uid,
        },
    }
    
    iter := r.client.Single().Query(ctx, stmt)
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

func (r *EarningsRepository) CreateEarning(ctx context.Context, req *pb.CreateEarningRequest) (*pb.CreateEarningResponse, error) {
    mutation := spanner.Insert(
        "Earnings",
        []string{"uid", "earning_id", "amount", "date"},
        []interface{}{req.Uid, req.Earning.EarningId, req.Earning.Amount, req.Earning.Date.AsTime()},
    )

    _, err := r.client.Apply(ctx, []*spanner.Mutation{mutation})
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to insert earning: %v", err)
    }

    return &pb.CreateEarningResponse{Status: "SUCCESS"}, nil
}
