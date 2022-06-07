package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

import (
	"context"
	"s_grpc/service"
)

func main(){
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalln("服务端出错或者未启动")
	}
	defer conn.Close()
	fmt.Println("xxx")
	userClient := service.NewUserServerClient(conn)

	//res, err := userClient.GetUserMessage(context.Background(), request)
	//if err != nil{
	//	log.Fatalln("查询库存出错")
	//}
	fmt.Println("aaa")

	// 添加用户
	request := &service.User{
		Username: "kebi",
		Age: 0,
	}
	res, err := userClient.AddUserMessage(context.Background(), request)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println("插入的name和age：")
	fmt.Println("name:", res.Username)
	fmt.Println("age:", res.Age)
	q := &service.UserRequest{
		Username: "kebi",
	}
	res1, err := userClient.GetUserMessage(context.Background(),q)
	fmt.Println("得到的name和age：")
	fmt.Println("name:", res1.Username)
	fmt.Println("age:", res1.Age)
}