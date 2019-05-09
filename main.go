package main

import (
	"fmt"
	"github.com/guapo-organizations/go-micro-secret/frame_tool"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "zldz/account-service/proto/account"
	myservice "zldz/account-service/service"
)

func main() {

	//加载一些公共的组件，比如数据库之类的
	my_frame_tool := new(frame_tool.LyFrameTool)
	my_frame_tool.ConfigPath = "./config"
	port := my_frame_tool.Run()

	//开始监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("account服务监听失败: %v", err)
	}

	//构建服务
	s := grpc.NewServer()
	//注册服务
	pb.RegisterAccountServer(s, &myservice.AccountService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务出现异常: %v", err)
	}
}
