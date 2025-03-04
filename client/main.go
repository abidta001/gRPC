package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "admin/admin/proto" // Import the generated package

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Aabi"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Println("Server Response:", resp.Message)
}
