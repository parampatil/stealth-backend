package repository

import (
    "context"
    pb "github.com/parampatil/stealth-backend/proto"
)

type HealthCheckRepository struct{}

func NewHealthCheckRepository() *HealthCheckRepository {
    return &HealthCheckRepository{}
}

func (r *HealthCheckRepository) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
    return &pb.HealthCheckResponse{
        Status: "SERVING",
    }, nil
}
