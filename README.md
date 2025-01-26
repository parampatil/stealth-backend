# stealth-backend

## Project setup
### Initialize Go module
```bash
go mod init github.com/stealth/backend
```
### Install Dependencies
```bash
go get google.golang.org/grpc
go get google.golang.org/protbuf
go get cloud.google.com/go/spanner/...
```

### Build proto file
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/health.proto
```

### Run the project
```bash
go run cmd/main.go
```