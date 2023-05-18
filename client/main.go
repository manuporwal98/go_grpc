package main

import (
	"log"

	pb "github.com/manuporwal98/go_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8083" //same port between server and client for their interaction
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() //connection closing,, after everthing is over, close the connection between server and client , defer ensures the close funtion executes at the very end
	client := pb.NewGreetServiceClient(conn)
	//names := &pb.NamesList{
	//	Names:= []string{"Manu","Chinu","Avi"},
	//}
	CallSayHello(client)
	//basically we are trying to implement something which helps client to access the function directly.
}
