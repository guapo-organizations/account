package test

import (
	"context"
	"flag"
	pb "github.com/guapo-organizations/account-service/proto/account"
	"github.com/guapo-organizations/go-micro-secret/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
	"time"
)

//tls ssl 加密测试
func TestTlsSSl(t *testing.T) {

	var email string
	flag.StringVar(&email, "email", "", "验证码邮件收件人")
	flag.Parse()
	// Set up a connection to the server.
	//tls配置,文件好像是通过第二个参数也就是 x.test.youtube.com生成的...fuck！！！
	creds, err := credentials.NewClientTLSFromFile(tls.Path("ca.pem"), "zldz.com")
	//连接的时候添加tls配置，公钥？不懂
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendEmailCode(ctx, &pb.EmailOrPhoneRequest{Email: email})
	if err != nil {
		t.Fatalf("rpc发生错误: %v", err)
	}

	t.Logf("结果:%v", r)
}
