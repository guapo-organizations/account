package test

import (
	"github.com/guapo-organizations/go-micro-secret/consul"
	"github.com/guapo-organizations/go-micro-secret/frame_tool"
	"github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"testing"
)

//测试通过consul服务发现找到服务地址
func TestFindConsulService(t *testing.T) {
	//加载一些公共的组件，比如数据库之类的
	my_frame_tool := frame_tool.LyFrameTool{
		ConfigPath: "../config",
	}
	my_frame_tool.Run()

	service_info := service.GetGrpcServiceInfo()
	config := consul.CreateConfig("localhost", "8500")
	service_id := consul.CreateAgentServiceUniqueID(service_info.Name, service_info.Ip, service_info.Port)
	agent_service_info, err := consul.FindService(config, service_id)
	if err != nil {
		t.Fatal("找不到服务", err)
	}
	t.Log("找到服务了", agent_service_info)
}
