package fibonacci

import (
	"context"
	"fmt"
	"pkg/fibonacci_api/pkg/fibonacci_api/pkg/fibonacci_api"
)

type GRPCServer struct {
}

func createFibonacciInterval(x, y int64) []int64 {
	slice := make([]int64, y)
	var i int64
	for i = 0; i <= y; i++ {
		if i == 0 || i == 1 {
			slice[i] = 0
		} else {
			slice[i] = slice[i-1] + slice[i-2]
		}
	}
	fmt.Println(slice)
	return slice[x:]
}

func (s *GRPCServer) Fibonacci(ctx context.Context, req *fibonacci_api.FibonacciRequest) (*fibonacci_api.FibonacciResponse, error) {
	x := req.GetX()
	y := req.GetY()
	return &fibonacci_api.FibonacciResponse{Slice: createFibonacciInterval(x, y)}, nil
}
