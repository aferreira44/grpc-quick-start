package servers

import (
	"context"
	"time"

	"github.com/aferreira44/grpc-quick-start/pb"
)

// Math is sfkglakld lkslkgjasldkfj kasjdlfk jlaskjdflk jaskdj
type Math struct {
}

// Sum is ...
func (m *Math) Sum(ctx context.Context, req *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	res := req.Sum.A + req.Sum.B

	return &pb.NewSumResponse{
		Result: res,
	}, nil
}

// Fibonacci is
func (m *Math) Fibonacci(in *pb.FibonacciRequest, stream pb.MathService_FibonacciServer) error {
	n := in.GetN()

	var i int32

	for i = 1; i <= n; i++ {
		res := &pb.FibonacciResponse{
			Result: FibonacciResolver(i),
		}
		stream.Send(res)
		time.Sleep(time.Second * 1 / 2)
	}
	return nil
}

// FibonacciResolver is ...
func FibonacciResolver(n int32) int32 {
	if n <= 1 {
		return n
	}

	return FibonacciResolver(n-1) + FibonacciResolver(n-2)
}
