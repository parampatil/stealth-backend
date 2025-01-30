package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "cloud.google.com/go/spanner"
    "github.com/parampatil/stealth-backend/internal/server"
)

func main() {
    ctx := context.Background()

    // Spanner client Initialization
    spannerClient, err := spanner.NewClient(ctx, "projects/stealth-448823/instances/stealth-backend-db/databases/youtube_earnings")
    if err != nil {
        log.Fatalf("Failed to create Spanner client: %v", err)
    }

    // Server Initialization
    srv := server.NewServer(spannerClient)
    
    // Handle graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    // Start server in goroutine
    go func() {
        if err := srv.Start(":8080"); err != nil {
            log.Fatalf("Failed to start server: %v", err)
        }
    }()

    <-sigChan
    log.Println("Shutting down server...")
    srv.Stop()
}
