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

func (s *HealthServer) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
    message := "Hello, " + req.Name + "!"
    return &pb.GreetingResponse{
        Message: message,
    }, nil
}