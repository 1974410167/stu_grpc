package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"s_grpc/service"
)


func main(){
	conn, err := grpc.Dial(":8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalln("服务端出错或者未启动")
	}
	defer conn.Close()
	manageClient := service.NewManageServiceClient(conn)

	fmt.Println("aaa")

	// 查找用户
	request := &service.ManageRequest{
		Username: "kebi",
	}
	res, err := manageClient.GetUserMessage(context.Background(), request)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("调用成功")
	fmt.Println("name:", res.Username)
	fmt.Println("age:", res.Age)

}