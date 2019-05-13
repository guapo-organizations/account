package main

import (
	"fmt"
	pb "github.com/guapo-organizations/account-service/proto/account"
	myservice "github.com/guapo-organizations/account-service/service"
	"github.com/guapo-organizations/go-micro-secret/frame_tool"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	//加载一些公共的组件，比如数据库之类的
	my_frame_tool := frame_tool.LyFrameTool{
		ConfigPath: "./config",
	}

	//返回的是grpc服务的链接信息
	grpc_service_info := my_frame_tool.Run()

	//开始监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpc_service_info.Port))
	if err != nil {
		log.Fatalf("account服务监听失败: %v", err)
	}

	log.Println(fmt.Sprintf("服务端口:%s", grpc_service_info.Port))

	//构建服务
	s := grpc.NewServer()
	//注册服务
	pb.RegisterAccountServer(s, &myservice.AccountService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务出现异常: %v", err)
	}
}
