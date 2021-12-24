package main

import (
	"fmt"
	"log"

	"github.com/yasensim/grpc_tests/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hellow from Client")
	client_connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could nor connect: %v", err)
	}
	defer client_connection.Close()

	c := greetpb.NewGreetServiceClient(client_connection)
	fmt.Printf("Created client: %f", c)
}
