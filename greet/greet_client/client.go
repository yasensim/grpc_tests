package main

import (
	"context"
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
	//	fmt.Printf("Created client: %f", c)
	doUnary(c)
}
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary gRPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Yasen",
			LastName:  "Simeonov",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)

}
