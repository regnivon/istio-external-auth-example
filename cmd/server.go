package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"test/pkg/protos"

	"google.golang.org/grpc"
)

type HelloService struct {}


func (s *HelloService) SayHello(ctx context.Context, req *protos.SayHelloRequest) (*protos.SayHelloResponse, error) {
	fmt.Println(req.Hi)

	return &protos.SayHelloResponse{
		Successful: true,
	}, nil
}

func main() {
	serv := grpc.NewServer()
	hello := &HelloService{}
	protos.RegisterHelloServiceServer(serv, hello)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}