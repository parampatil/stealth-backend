package repository

import (
    "context"
    "fmt"
    pb "github.com/parampatil/stealth-backend/proto"
)

type GreetingRepository struct{}

func NewGreetingRepository() *GreetingRepository {
    return &GreetingRepository{}
}

func (r *GreetingRepository) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
    message := fmt.Sprintf("Hello, %s!", req.Name)
    return &pb.GreetingResponse{
        Message: message,
    }, nil
}
