package main

import (
	"context"
	"flag"
	pb "github.com/guapo-organizations/account-service/proto/account"
	"github.com/guapo-organizations/go-micro-secret/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

var email string

func init() {
	flag.StringVar(&email, "email", "", "验证码邮件收件人")
	flag.Parse()
}
func main() {
	// Set up a connection to the server.
	//tls配置,文件好像是通过第二个参数也就是 x.test.youtube.com生成的...fuck！！！
	creds, err := credentials.NewClientTLSFromFile(tls.Path("ca.pem"), "x.test.youtube.com")
	//连接的时候添加tls配置，公钥？不懂
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendEmailCode(ctx, &pb.EmailOrPhoneRequest{Email: email})
	if err != nil {
		log.Fatal("rpc发生错误: %v", err)
	}

	log.Printf("结果:%v", r)
}
