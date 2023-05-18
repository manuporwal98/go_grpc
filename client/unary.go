package main

import (
	"context"
	"log"
	"time"

	pb "github.com/manuporwal98/go_grpc/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	//https://golang.cafe/blog/golang-context-with-timeout-example.html , documentation on timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() //will be executed in the very end for cleanup
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
