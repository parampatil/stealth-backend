package main

import (
	"context"
	"log"
	"net"

	"cloud.google.com/go/spanner"
	"github.com/parampatil/stealth-backend/internal/server"
	pb "github.com/parampatil/stealth-backend/proto"
	"google.golang.org/grpc"
)

// TODO: Implement the main function in goroutine
// TODO: Write the project id in environment variable
func main() {
    ctx := context.Background()

    spannerClient, err := spanner.NewClient(ctx, "projects/stealth-448823/instances/stealth-backend-db/databases/youtube_earnings")
    if err != nil {
        log.Fatalf("Failed to create Spanner client: %v", err)
    }
    defer spannerClient.Close()

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    healthServer := server.NewHealthServer(spannerClient)
    pb.RegisterHealthServiceServer(s, healthServer)

    log.Printf("Server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}