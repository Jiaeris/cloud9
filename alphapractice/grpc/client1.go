package grpc

import (
	"google.golang.org/grpc"
	"context"
	"fmt"
	"strconv"
	//"time"
	"sync"
	"time"
)

const (
	address = "localhost:10086"
	times   = 100
)

func openClient() {

	//模拟100个客户端并发请求
	i := 0
	wg := sync.WaitGroup{}
	for i < times {
		wg.Add(1)
		go func(i int) {
			client, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				fmt.Println(err)
			}
			defer client.Close()
			cc := NewControllerClient(client)
			time.Sleep(time.Second)
			res, err := cc.SayHello(context.Background(), &HelloRequest{Name: "Yunga" + strconv.Itoa(i)}, grpc.FailFast(true))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("clent1: ", res.GetMessage())

			streamClient, err := cc.SayStream(context.Background(), grpc.FailFast(true))
			req := &RequestStreamObj{
				Requestid: int32(i),
				Data:      "request data",
			}
			streamClient.Send(req)

			wg.Done()
		}(i)
		i++
	}
	wg.Wait()
	fmt.Println("clients request finshed.")
}
