//package grpc

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	hellopb "github.com/mccomack/goback/proto"
)

type greeterServer struct {
	hellopb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(
	ctx context.Context,
	req *hellopb.HelloRequest,
) (*hellopb.HelloResponse, error) {
	log.Printf("name: %s", req.GetName())

	return &hellopb.HelloResponse{
		Message: "Hello, " + req.GetName(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	hellopb.RegisterGreeterServer(
		grpcServer,
		&greeterServer{},
	)

	log.Println("gRPC server listening on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//grpcServer.GracefulStop()
}
