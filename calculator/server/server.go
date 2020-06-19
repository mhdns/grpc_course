package main

import (
	"context"
	"fmt"
	"grpc_course/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	v1, v2 := req.GetValue().Value1, req.Value.GetValue2()
	res := &calculatorpb.SumResponse{
		Sum: v1 + v2,
	}
	return res, nil
}

func (*server) PrimeNum(req *calculatorpb.PrimeNumRequest, stream calculatorpb.SumService_PrimeNumServer) error {
	fmt.Println("prime called...")
	number := req.GetValue()
	var n int32 = 2
	for number > 1 {
		if number%n == 0 {
			err := stream.Send(&calculatorpb.PrimeNumResponse{Value: n})
			if err != nil {
				log.Fatalf("Not able to send response")
			}
			number /= n
			continue
		}
		n++
	}
	return nil
}

func main() {
	fmt.Println("Hello world")

	li, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(li); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
