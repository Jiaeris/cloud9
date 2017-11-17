package grpc

import (
	"os/signal"
	"os"
	"fmt"
)

//生成grpc proto
//protoc --go_out=plugins=grpc:. ./rpcmsg.proto

func test() {
	//开启服务端
	go openServer()
	//开启客户端
	go openClient()
	//以上为普通调用，权限？身份验证？

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			break
		}
	}
}
