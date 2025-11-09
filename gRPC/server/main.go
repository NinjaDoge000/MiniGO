package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gRPC/proto/greeter"

	"google.golang.org/grpc"
)

// Server struct (must embed UnimplementedGreeterServer)
type server struct {
	pb.UnimplementedGreeterServer
}

// Implement the SayHello method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request for name: %s", req.Name)
	return &pb.HelloReply{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("🚀 gRPC server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
