package handler

import (
	"context"
	pb "grpc/protob/example"
)

type GreeterService struct{}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}
