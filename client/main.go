package main

import (
	"log"

	pb "github.com/vaibhavvvvv/grpc0/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	//for client to connect with server :
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn't connect : %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList{
	// 	Names: []string{"Vaibhav", "Alice", "Bob"},
	// }
	callSayHello(client)
}
