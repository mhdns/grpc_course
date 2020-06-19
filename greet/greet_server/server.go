package main

import (
	"context"
	"fmt"
	"grpc_course/greet/greetpb"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fname := req.GetGreeting().GetFirstName()
	result := "Hello " + fname
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetMany(req *greetpb.GreetManyRequest, stream greetpb.GreetService_GreetManyServer) error {
	fname := req.GetGreeting().GetFirstName()
	for i := 1; i < 11; i++ {
		result := "Hello " + fname + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	fmt.Println("Hello world")

	li, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(li); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
