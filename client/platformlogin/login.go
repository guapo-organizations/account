package main

import (
	"context"
	"fmt"
	"github.com/guapo-organizations/account-service/proto/account"
	"github.com/guapo-organizations/sms-service/tls"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	//tls 和 ssl加密
	creds, err := tls.GetClientTLSFromFile("ca.pem","zldz.com")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", "50051"), grpc.WithTransportCredentials(creds))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := account.NewAccountClient(conn)

	response, err := client.OpenPlaformLogin(ctx, &account.OpenPlaformLoginRequest{
		Name:       "梁宇",
		TypeId:     1,
		PlatformId: "13123asd",
		UnionId:    "asdfdd",
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
