package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDiStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Client Streaming gRPC")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Yasen",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Petya",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Marto",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Vladi",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}
	// interate over the slice and send each message individualy
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reveiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do BiDir Streaming gRPC")
	// create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while create a stream: %v", err)
	}
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Yasen",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Petya",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Marto",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Vladi",
			},
		},
	}
	// create wait channel to block bellow
	waitc := make(chan struct{})
	// send a bunch of meessages (go routine)
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// Receive a bunch of meessages (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving a stream: %v", err)
				break
			}
			fmt.Printf("Recieved: %v \n ", res.GetResult())
			//			time.Sleep(3000 * time.Millisecond)
		}
		close(waitc)
	}()
	// Block until everything is done
	<-waitc

}
