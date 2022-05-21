package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/service"
	"log"
	"net"
)



func main(){
	rpcServer := grpc.NewServer()

	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil{
		log.Fatal("启动监听出错", err)
	}
	rpcServer.Serve(listener)
	fmt.Println("启动成功")
}