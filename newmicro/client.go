package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	imooc "learn_docker/newmicro/proto/cap"
)

func main() {
	//实例化
	service := micro.NewService(
		micro.Name("cap.imooc.client"),
	)
	//初始化
	service.Init()

	capImooc := imooc.NewCapService("cap.imooc.server", service.Client())
	res, err := capImooc.SayHello(context.TODO(), &imooc.SayRequest{Message: "跟着Cap老师学习微服务！"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Answer)
}
