//package grpc

package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	hellopb "github.com/mccomack/goback/proto"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := hellopb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	response, err := client.SayHello(ctx, &hellopb.HelloRequest{
		Name: "Alice",
	})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Println(response.GetMessage())
}
