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
	service.RegisterManageServiceServer(rpcServer, service.ManagerServiceIns)
	listener, err := net.Listen("tcp", ":8003")
	if err != nil{
		log.Fatal("启动监听出错", err)
	}
	rpcServer.Serve(listener)
	fmt.Println("启动成功")
}
