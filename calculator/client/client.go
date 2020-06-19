package main

import (
	"context"
	"fmt"
	"grpc_course/calculator/calculatorpb"
	"io"
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

	// fmt.Println(sum(conn, 5, 6))

	results := make(chan int32)

	go primeNum(conn, 1847198980, results)

	for v := range results {
		// time.Sleep(time.Second)
		fmt.Println(v)
	}

}

func primeNum(conn *grpc.ClientConn, v1 int32, results chan int32) {
	c := calculatorpb.NewSumServiceClient(conn)

	req := &calculatorpb.PrimeNumRequest{Value: v1}

	streamRes, err := c.PrimeNum(context.Background(), req)
	if err != nil {
		log.Fatalf("No able to recieve response")
	}

	for {
		msg, err := streamRes.Recv()
		if err == io.EOF {
			fmt.Println("end of message")
			close(results)
			break
		} else if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		results <- msg.GetValue()
	}
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
