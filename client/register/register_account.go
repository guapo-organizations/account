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

func main(){
	//tls 和 ssl加密
	creds, err := tls.GetClientTLSFromFile("ca.pem","zldz.com")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", "50051"), grpc.WithTransportCredentials(creds))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := account.NewAccountClient(conn)

	response, err := client.RegisterAccountByEmail(ctx, &account.RegisterAccountEmailRequester{
		Email:  "51785816@qq.com",
		Code:   "oh09G3",
		Passwd: "123456789",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(response)
}
