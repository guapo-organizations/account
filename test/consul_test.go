package test

import (
	"github.com/guapo-organizations/go-micro-secret/consul"
	"log"
	"testing"
)

//测试通过consul服务发现找到服务地址
func TestFindConsulService(t *testing.T) {
	config := consul.CreateConfig("localhost", "8500")
	//tag传中文会报错
	agent_service_info, err := consul.FindService(config, "account", "account")
	if err != nil {
		t.Fatal("找不到服务", err)
	}
	log.Printf("id:%s,端口:%d,tags:%s",agent_service_info.Address,agent_service_info.Port,agent_service_info.Tags)
}
