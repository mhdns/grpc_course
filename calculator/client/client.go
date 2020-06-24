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

	// Unary
	// fmt.Println(sum(conn, 5, 6))

	// Server Streaming
	// results := make(chan int32)

	// go primeNum(conn, 1847198980, results)

	// for v := range results {
	// 	// time.Sleep(time.Second)
	// 	fmt.Println(v)
	// }

	// Client streaming
	fmt.Println(averageAge(conn, []int32{31, 42, 12, 19}))

}

func averageAge(conn *grpc.ClientConn, ages []int32) (int32, error) {
	c := calculatorpb.NewSumServiceClient(conn)

	stream, err := c.AverageAge(context.Background())
	if err != nil {
		return 0, nil
	}

	for _, v := range ages {
		stream.Send(&calculatorpb.AverageAgeRequest{Value: v})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while recieving response...")
	}

	return res.GetValue(), nil
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
