package server

import (
    "context"
    pb "github.com/parampatil/stealth-backend/proto"
)

type HealthServer struct {
    pb.UnimplementedHealthServiceServer
}

func NewHealthServer() *HealthServer {
    return &HealthServer{}
}

func (s *HealthServer) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
    return &pb.HealthCheckResponse{
        Status: "SERVING",
    }, nil
}
