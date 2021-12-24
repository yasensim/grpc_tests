package main

import (
	"context"
	"fmt"
	"io"
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
	//	doUnary(c)
	doServerStreaming(c)
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
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Server Streaming gRPC")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Yasen",
			LastName:  "Simeonov",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Server Streaming RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we have reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("response from GreetManyTimes: %v", msg.GetResult())
	}

}
