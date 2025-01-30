package controller

import (
    "cloud.google.com/go/spanner"
    "context"
    pb "github.com/parampatil/stealth-backend/proto"
    "github.com/parampatil/stealth-backend/internal/repository"
)

type HealthController struct {
    pb.UnimplementedHealthServiceServer
    healthCheckRepo  *repository.HealthCheckRepository
    greetingRepo    *repository.GreetingRepository
    earningsRepo    *repository.EarningsRepository
    demoDataRepo    *repository.DemoDataRepository
}

func NewHealthController(client *spanner.Client) *HealthController {
    return &HealthController{
        healthCheckRepo:  repository.NewHealthCheckRepository(),
        greetingRepo:    repository.NewGreetingRepository(),
        earningsRepo:    repository.NewEarningsRepository(client),
        demoDataRepo:    repository.NewDemoDataRepository(client),
    }
}

func (c *HealthController) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
    return c.healthCheckRepo.Check(ctx, req)
}

func (c *HealthController) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
    return c.greetingRepo.Greet(ctx, req)
}

func (c *HealthController) GetEarnings(ctx context.Context, req *pb.GetEarningsRequest) (*pb.GetEarningsResponse, error) {
    return c.earningsRepo.GetEarnings(ctx, req)
}

func (c *HealthController) GenerateDemoData(ctx context.Context, req *pb.GenerateDemoDataRequest) (*pb.GenerateDemoDataResponse, error) {
    return c.demoDataRepo.GenerateDemoData(ctx, req)
}
