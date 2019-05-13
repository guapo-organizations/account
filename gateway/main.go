package main

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/guapo-organizations/go-micro-secret/frame_tool"
	"google.golang.org/grpc"
	"net/http"

	//grpc网关
	gw "github.com/guapo-organizations/account-service/proto/account"
	grpc_gateway_service_info "github.com/guapo-organizations/go-micro-secret/frame_tool/service"
)

//grpc网关

func run() error {
	my_frame_tool := frame_tool.LyFrameTool{
		ConfigPath: "./config",
	}
	my_frame_tool.Run()
	gate_way_service_info := grpc_gateway_service_info.GetGrpcGateWayServiceInfo()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterAccountHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", gate_way_service_info.Ip, gate_way_service_info.Port), opts)
	if err != nil {
		return err
	}

	//grpc网关的服务端监听,也就是开启Web服务
	return http.ListenAndServe(fmt.Sprintf(":%s", gate_way_service_info.GatewayPort), mux)
}
func main() {
	//不知道这个是什么，github上是这样写的，照样抄就对了,应该是个日志组件之类的
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
