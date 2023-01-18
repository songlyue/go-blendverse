package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NacosInit() {
	clientConfig := constant.ClientConfig{
		NamespaceId: "f91502c6-677f-4bed-8560-a1d266a34fea",
		Username:    "nacos",
		Password:    "zeewain@123",
		LogDir:      "./logs",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.0.58",
			Port:   38848,
		}}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		_ = fmt.Errorf("创建服务发现客户端失败")
	}

	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.3.248",
		Port:        28888,
		ServiceName: "mock-fusion-server",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"context-path": "blend", "strip-prefix": "true", "is-jwt": "false"},
	})
	if !success {
		_ = fmt.Errorf("服务注册上nacos失败")
	}
}
