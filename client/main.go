package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "github.com/guapo-organizations/account-service/proto/account"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RegisterAccountByEmailOrPhone(ctx, &pb.EmailOrPhoneRequest{Phone: "13647897503"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("结果:%v", r)
}
