package grpc

import (
	"net"
	"google.golang.org/grpc"
	"context"
	"fmt"
	"io"
)

const (
	port = ":10086"
)

type data int

func openServer() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	var x data
	RegisterControllerServer(s, &x)
	if err = s.Serve(l); err != nil {
		panic(err)
	}
}

func (d *data) SayHello(ctx context.Context, req *HelloRequest) (res *HelloResponse, err error) {
	reqMsg := req.GetName()
	fmt.Println("server: ", reqMsg)
	res = &HelloResponse{
		Message: "hello " + reqMsg,
	}
	return res, err
}

func (d *data) SayStream(stream *controllerSayStreamServer) error {

	req, err := stream.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
