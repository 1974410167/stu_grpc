package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"s_grpc/service"
)

func main(){
	rpcServer := grpc.NewServer()
	service.RegisterUserServerServer(rpcServer, service.UserServiceIns)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil{
		log.Fatal("启动监听出错", err)
	}
	rpcServer.Serve(listener)
	fmt.Println("启动成功")
}
