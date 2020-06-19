package main

import (
	"context"
	"fmt"
	"grpc_course/greet/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connedt: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	// doUnary(c)
	doServerStream(c)
}

func doServerStream(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetManyRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Anas",
			LastName:  "Maricar",
		},
	}

	streamRes, err := c.GreetMany(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not get response: %v", err)
	}
	for {
		msg, err := streamRes.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response: %v", msg.GetResult())
	}

}

func doUnary(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Anas",
			LastName:  "Maricar",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not get response: %v", err)
	}
	fmt.Println(res.Result)
}
