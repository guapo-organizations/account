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
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendEmailCode(ctx, &pb.EmailOrPhoneRequest{Email: "876499812@qq.com"})
	if err != nil {
		log.Fatal("rpc发生错误: %v", err)
	}

	log.Printf("结果:%v", r)
}
