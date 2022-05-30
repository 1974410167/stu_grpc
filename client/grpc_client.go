package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/service"
	"log"
)

func main(){
	conn, err := grpc.Dial(":8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalln("服务端出错或者未启动")
	}
	defer conn.Close()
	request := &service.ProductRequest{
		ProdId: 123,
		User: &service.User{
			Username: "GEHAOYUAN",
			Age: 22,
			Addresses: []string{"zhengzhou", "donglu"},
		},
	}

	prodClient := service.NewProdServiceClient(conn)
	res, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil{
		log.Fatalln("查询库存出错")
	}
	fmt.Println("成功", res.ProdStock)
	fmt.Println("成功", res.Message)

}