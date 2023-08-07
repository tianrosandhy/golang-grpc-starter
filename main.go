package main

import (
	"fmt"
	"log"
	"net"

	"grpc/handler"
	"grpc/lib"
	appointmentPB "grpc/protob/appointment"
	examplePB "grpc/protob/example"
	"grpc/services"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var Environment = map[string]interface{}{
	"port":    50051,
	"db_user": "postgres",
	"db_pass": "postgres",
	"db_host": "localhost",
	"db_port": 5432,
	"db_name": "postgres",
}

func init() {
	lib.LoadEnvironment(Environment)
	services.InitDatabase()
}

func main() {
	// Start the gRPC server
	grpcServer := grpc.NewServer()
	examplePB.RegisterGreeterServer(grpcServer, &handler.GreeterService{})
	appointmentPB.RegisterAppointmentServer(grpcServer, &handler.AppointmentService{})

	port := viper.GetInt("PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server listening on :%d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
