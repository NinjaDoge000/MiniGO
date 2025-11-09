package main

import (
	"context"
	"log"
	"time"

	pb "gRPC/proto/greeter"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Laxman"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Server response: %s", resp.Message)
}
