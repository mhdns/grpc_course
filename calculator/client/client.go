package main

import (
	"context"
	"fmt"
	"grpc_course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client says Hello")

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connedt: %v", err)
	}
	defer conn.Close()

	fmt.Println(sum(conn, 5, 6))

}

func sum(conn *grpc.ClientConn, v1, v2 int32) (int32, error) {
	c := calculatorpb.NewSumServiceClient(conn)

	req := &calculatorpb.SumRequest{
		Value: &calculatorpb.Values{
			Value1: v1,
			Value2: v2,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		return 0, err
	}

	return res.Sum, nil
}
