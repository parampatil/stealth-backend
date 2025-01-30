package repository

import (
    "context"
    "fmt"
    "math/rand"
    "time"
    "cloud.google.com/go/spanner"
    pb "github.com/parampatil/stealth-backend/proto"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type DemoDataRepository struct {
    client *spanner.Client
}

func NewDemoDataRepository(client *spanner.Client) *DemoDataRepository {
    return &DemoDataRepository{client: client}
}

func (r *DemoDataRepository) GenerateDemoData(ctx context.Context, req *pb.GenerateDemoDataRequest) (*pb.GenerateDemoDataResponse, error) {
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

    _, err := r.client.Apply(ctx, mutations)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to generate demo data: %v", err)
    }

    return &pb.GenerateDemoDataResponse{Status: "SUCCESS"}, nil
}
