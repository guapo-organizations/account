package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"

	//grpc网关
	gw "github.com/guapo-organizations/account-service/proto/account"
)

//grpc网关

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterAccountHandlerFromEndpoint(ctx, mux, "106.12.76.73:5051", opts)
	if err != nil {
		return err
	}

	//grpc网关的服务端监听,也就是开启Web服务
	return http.ListenAndServe(":8080", mux)
}
func main() {
	//不知道这个是什么，github上是这样写的，照样抄就对了,应该是个日志组件之类的
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
