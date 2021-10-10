package main

import (
	"context"
	"flag"
	"fmt"
	"go-advance/chapter1/grpchealth/api"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"time"
)

var (
	port = flag.Int("prot", 9000, "the port to serve on")

	// 双方的一个约定
	service = "hello"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf(":%v", *port), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()
	// 新建一个客户端
	c := api.NewGreeterClient(conn)
	healthcheck := healthpb.NewHealthClient(conn)
	// 调用服务端函数
	for {
		// 发送前先进行健康检查
		resp, err := healthcheck.Check(context.Background(),
			&healthpb.HealthCheckRequest{
				Service: service,
			},
		)
		if err != nil {
			panic(err)
		}
		switch resp.Status {
		case healthpb.HealthCheckResponse_SERVING:
			// 如果健康检查成功则返回
			r, err := c.SayHello(context.Background(), &api.HelloRequest{Name: "horika"})
			if err != nil {
				fmt.Printf("调用服务端代码失败: %s", err)
				return
			}
			fmt.Printf("调用成功: %s", r.Message)
		default:
			fmt.Println("健康检查失败，服务不可用")
		}
		time.Sleep(1 * time.Second)
	}
}
