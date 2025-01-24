package main

import (
    "log"
    "net"
    "github.com/parampatil/stealth-backend/internal/server"
    pb "github.com/parampatil/stealth-backend/proto"
    "google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    healthServer := server.NewHealthServer()
    pb.RegisterHealthServiceServer(s, healthServer)

    log.Printf("Server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
