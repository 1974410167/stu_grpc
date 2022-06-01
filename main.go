package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

//详见 https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md

func main(){

	// nacos客户端的配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// nacos服务端的配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}


	// 根据配置生成动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	HandlerError(err)
	// 发布配置
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  "dataId",
		Group:   "group",
		Content: "hello world!222222"},
	)
	HandlerError(err)
	fmt.Println(success)
	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "nacos_test",
		Group:  "DEFAULT_GROUP"})
	HandlerError(err)
	fmt.Println("获取配置：", content)
	// 监听配置变化
	configClient.ListenConfig(vo.ConfigParam{
		DataId: "nacos_test",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	//HandlerError(err)


	// 服务发现客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	// 服务注册
	success, err = namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",   // grpc服务IP
		Port:        8000,          // grpc服务端口
		ServiceName: "demo1.go",   // 给grpc服务起个名字
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc":"shanghai"},
	})
	HandlerError(err)
	fmt.Println("服务注册1:", success)

	success, err = namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",   // grpc服务IP
		Port:        8001,          // grpc服务端口
		ServiceName: "demo2.go",   // 给grpc服务起个名字
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc":"shanghai"},
	})
	HandlerError(err)
	fmt.Println("服务注册2:", success)

	// 获取服务信息
	services, err := namingClient.GetService(vo.GetServiceParam{
		ServiceName: "demo1.go",
	})
	HandlerError(err)
	fmt.Println("hosts:", services.Hosts)
	fmt.Println("name:", services.Name)
	fmt.Println("clusters:", services.Clusters)
	// 获取服务的实例列表
	instances, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: "demo1.go",
	})
	HandlerError(err)
	fmt.Println("instance:", instances)
	// 随机获取一个实例
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "demo1.go",
	})
	HandlerError(err)
	fmt.Println("随机获取一个实例1：", instance)
	// 随机获取一个实例
	instance1, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "demo2.go",
	})
	HandlerError(err)
	fmt.Println("随机获取一个实例2：", instance1)
	fmt.Println("实例地址:", instance1.Ip)
	fmt.Println("实例端口:", instance1.Port)
	fmt.Println("实例名字:", instance1.ServiceName)

	// 获取服务列表
	serviceInfos, err := namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: "",
		PageNo:   1,
		PageSize: 10,
	})
	fmt.Println("serviceInfos:", serviceInfos)
}

func HandlerError(v any){
	if v != nil{
		panic(v)
	}
}



