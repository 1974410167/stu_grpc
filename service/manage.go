package service

import (
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)



var ManagerServiceIns = &ManagerService{}


type ManagerService struct {
}

func (m *ManagerService) mustEmbedUnimplementedManageServiceServer(){

}

// 实现user.proto定义的方法，根据username拿到用户信息
func (m *ManagerService) GetUserMessage(c context.Context,request *ManageRequest) (*User, error) {
	name := request.Username
	// 获得服务发现客户端
	q ,err:= GetServerFindClient()
	HandlerError(err)
	// 从服务发现客户端根据服务名得到一个实例
	instance1, err := q.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "user_server.go",
	})
	if err != nil{
		fmt.Println(err)
	}
	ip := instance1.Ip
	port := instance1.Port
	tar := ip + ":" + string(port)
	// 连接实例，获取实例句柄
	conn, err := grpc.Dial(tar, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalln("MangerServer调用UserServer出错")
	}
	defer conn.Close()
	// 获得用户客户端
	userClient := NewUserServerClient(conn)
	t := &UserRequest{
		Username: name,
	}
	// 根据用户客户端调用客户方法
	res, err := userClient.GetUserMessage(context.Background(), t)
	if err != nil{
		log.Fatalln("MangerServer调用UserServer出错来查询用户信息出错")
	}
	return res, nil
}

func HandlerError(v interface{}){
	if v != nil{
		panic(v)
	}
}
