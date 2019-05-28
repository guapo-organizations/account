package main

import (
	"context"
	"fmt"
	"github.com/guapo-organizations/account-service/proto/account"
	"github.com/guapo-organizations/go-micro-secret/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

func main() {
	//tls 和 ssl加密
	creds, err := credentials.NewClientTLSFromFile(tls.Path("ca.pem"), "zldz.com")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", "50051"), grpc.WithTransportCredentials(creds))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := account.NewAccountClient(conn)

	response, err := client.LoginByEmailPassword(ctx, &account.LoginByPasswordRequest{
		Email:  "51785816@qq.com",
		Passwd: "123456",
	})
	if err != nil {
		log.Fatalln("服务端错：", err)
	}
	log.Println(response)

	log.Println("token测试")
	response_token, err := client.TokenDecode(ctx, &account.TokenDecodeRequest{
		Token: response.Token,
	})
	if err != nil {
		log.Fatalln("token服务出错: ", err)
	}
	log.Printf("%v", response_token)
	log.Printf("%v", response_token.UserBaseInfo)
	log.Printf("你的名字：%s", response_token.UserBaseInfo.Name);
	log.Println("你的名字：" + response_token.UserBaseInfo.Name);
	log.Printf("%v", response_token.UserOpenPlatformInfo)
}
