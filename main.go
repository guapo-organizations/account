package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "zldz/account-service/proto/account"
	"zldz/account-service/service"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	port = ":50051"
)

func main() {

	//开始监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("account服务监听失败: %v", err)
	}

	//构建服务
	s := grpc.NewServer()

	//注册服务
	pb.RegisterAccountServer(s, &service.AccountService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
