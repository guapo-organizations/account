package main

import (
	"context"
	pb "github.com/guapo-organizations/account-service/proto/account"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	mychan := make(chan bool)
	go func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewAccountClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		for i := 0; i < 1000000; i++ {
			r, err := c.RegisterAccountByEmailOrPhone(ctx, &pb.EmailOrPhoneRequest{Phone: "13647897503"})
			if err != nil {
				log.Printf("could not greet: %v", err)
			}

			log.Printf("结果:%v", r)
		}
	}()

	go func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		c := pb.NewAccountClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		for i := 0; i < 1000000; i++ {
			r, err := c.RegisterAccountByEmailOrPhone(ctx, &pb.EmailOrPhoneRequest{Phone: "13647897504"})
			if err != nil {
				log.Printf("could not greet: %v", err)
			}

			log.Printf("结果:%v", r)
		}
	}()

	<-mychan
}
