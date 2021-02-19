package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"test/pkg/protos"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("34.74.148.75:50051", grpc.WithInsecure())
	md := metadata.Pairs("x-ext-authz", "allow")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	if err != nil {
		panic(err)
	}
	client := protos.NewHelloServiceClient(conn)
	req := &protos.SayHelloRequest{Hi: "Whats up"}
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Successful)

}
