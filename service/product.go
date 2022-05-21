package service

import (
	"context"
	"fmt"
)

var ProductService = &productService{}


type productService struct {
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {
	//panic("implement me")
}

func (p *productService) GetProductStock(c context.Context, request *ProductRequest) (*ProductResponse, error){
	fmt.Println(request.User.Username)
	fmt.Println(request.User.Age)
	fmt.Println(request.User.Addresses)

	msg := request.User.Username + string(request.User.Age) + request.User.Addresses[0]
	return &ProductResponse{
		ProdStock: request.ProdId*10,
		Message: msg,
	}, nil
}
  