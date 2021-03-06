package main

import (
	"fmt"
	pb "github.com/guapo-organizations/account-service/proto/account"
	myservice "github.com/guapo-organizations/account-service/service"
	"github.com/guapo-organizations/go-micro-secret/frame_tool"
	grpc_service_info "github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {

	//加载一些公共的组件，比如数据库之类的
	my_frame_tool := frame_tool.LyFrameTool{
		ConfigPath: "../config",
	}

	my_frame_tool.Run()
	service_info := grpc_service_info.GetGrpcServiceInfo()

	//tls配置
	//第一个参数是公钥，第二个参数是私钥
	creds, err := credentials.NewServerTLSFromFile("../config/tls/server1.pem", "../config/tls/server1.key")
	if err != nil {
		log.Fatalln("tls配置失败:", err)
	}
	//开始监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", service_info.Port))
	if err != nil {
		log.Fatalf("account服务监听失败: %v", err)
	}

	log.Println(fmt.Sprintf("服务端口:%s", service_info.Port))

	//构建服务,并启动tls配置
	s := grpc.NewServer(grpc.Creds(creds))
	//注册服务
	pb.RegisterAccountServer(s, &myservice.AccountService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务出现异常: %v", err)
	}
}
