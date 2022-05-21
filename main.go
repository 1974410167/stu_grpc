package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"grpc/service"
)

func main(){
	user := &service.User{
		Username: "haoyuan",
		Age: 100,
	}

	// 序列化
	proData, err := proto.Marshal(user)
	HandleError(err)
	// 反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(proData, newUser)
	HandleError(err)
	fmt.Println(newUser)
}


func HandleError(err any){
	if err != nil{
	}
}
