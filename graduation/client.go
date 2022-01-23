package main

import (
	"context"
	pb "go_advanced/graduation/api/v1"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:10011"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	r, err := c.GetUserName(context.Background(), &pb.GetUserNameRequest{Uid: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r.Username)
}
