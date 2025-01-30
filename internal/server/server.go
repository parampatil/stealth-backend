package server

import (
    "log"
    "net"
    "cloud.google.com/go/spanner"
    "google.golang.org/grpc"
    pb "github.com/parampatil/stealth-backend/proto"
    "github.com/parampatil/stealth-backend/internal/controller"
)

type Server struct {
    grpcServer *grpc.Server
    spannerClient *spanner.Client
}

// NewServer creates a new Server instance
func NewServer(spannerClient *spanner.Client) *Server {
    return &Server{
        grpcServer: grpc.NewServer(),
        spannerClient: spannerClient,
    }
}

// Start starts the gRPC server. (This function is called in the goroutine, with port as argument)
func (s *Server) Start(port string) error {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        return err
    }

    healthController := controller.NewHealthController(s.spannerClient)
    pb.RegisterHealthServiceServer(s.grpcServer, healthController) // Register the HealthServiceServer from proto file

    log.Printf("Server listening at %v", lis.Addr())
    return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
    if s.grpcServer != nil {
        s.grpcServer.GracefulStop()
    }
    if s.spannerClient != nil {
        s.spannerClient.Close()
    }
}
