package main

import (
	"context"
	"flag"
	"fmt"
	pb "go-advance/chapter1/grpchealth/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"time"
)

var (
	port  = flag.Int("prot", 9000, "the port to serve on")
	sleep = flag.Duration("sleep", time.Second*5, "duration between changes in health")

	// 双方的一个约定
	service = "hello"
)

type helloServer struct {
}

func (h *helloServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("hello, %v\n", request.Name),
	}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(s, healthcheck)
	pb.RegisterGreeterServer(s, &helloServer{})

	go func() {
		// 不断的更换健康检查状态
		next := healthpb.HealthCheckResponse_SERVING

		for {
			healthcheck.SetServingStatus(service, next)
			fmt.Printf("现在的状态是, %v\n", next)

			if next == healthpb.HealthCheckResponse_SERVING {
				next = healthpb.HealthCheckResponse_NOT_SERVING
			} else {
				next = healthpb.HealthCheckResponse_SERVING
			}

			time.Sleep(*sleep)
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
