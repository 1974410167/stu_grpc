package service

import (
	"context"
)


var UserServiceIns = &UserService{}


type UserService struct {
	names map[string]*User
	sign bool
}

func (u *UserService) mustEmbedUnimplementedUserServerServer() {
	//panic("implement me")
}


// 实现user.proto定义的方法，根据username拿到用户信息
func (u *UserService) GetUserMessage(c context.Context, request *UserRequest) (*User, error){
	name := request.Username
	q, ok := u.names[name]
	if ok{
		return q, nil
	}
	return nil, nil
}

// 添加一个用户信息
func (u *UserService) AddUserMessage(c context.Context, request *User) (*User, error){
	if u.sign == false{
		u.names = make(map[string]*User)
		u.sign = true
	}
	u.names[request.Username] = request
	return request, nil
}


