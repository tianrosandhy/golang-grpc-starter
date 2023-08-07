# GRPC Service Go

This is a golang base service to handle gRPC request  

### Setup
```
go run .
```

### Renew or create proto struct
```
protoc -I. --go_out=plugins=grpc:./ your-filename.proto
```

### Implement proto struct
- Create new handler that represent the proto file
- Register the handler in main.go 