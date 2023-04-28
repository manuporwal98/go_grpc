package main

import (
	"log"
	"net"

	pb "github.com/manuporwal98/go_grpc/proto" //pb is basically the fils which we get from the two new files created.
	"google.golang.org/grpc"
)

const (
	port = ":8083"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
