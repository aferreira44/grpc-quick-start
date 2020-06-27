package main

import (
	"context"
	"io"
	"log"

	"github.com/aferreira44/grpc-quick-start/pb"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer connection.Close()
	client := pb.NewMathServiceClient(connection)

	// request := &pb.NewSumRequest{
	// 	Sum: &pb.Sum{
	// 		A: 11,
	// 		B: 12,
	// 	},
	// }

	// res, err := client.Sum(context.Background(), request)

	request := &pb.FibonacciRequest{
		N: 10,
	}

	responseStream, err := client.Fibonacci(context.Background(), request)

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("Fibonacci: %v", stream.GetResult())
	}
}
