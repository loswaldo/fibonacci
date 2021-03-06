package fibonacci

import (
	"context"
	"errors"
	"github.com/loswaldo/fibonacci/pkg/fibonacci_api"
	"log"
)

type GRPCServer struct {
}

func createFibonacciInterval(x, y int64) []uint64 {
	slice := make([]uint64, y+1)
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	var i uint64
	for i = 0; i <= uint64(y); i++ {
		if i == 0 || i == 1 {
			slice[i] = i
		} else {
			slice[i] = slice[i-1] + slice[i-2]
		}
	}
	return slice[x:]
}

func (s *GRPCServer) Fibonacci(ctx context.Context, req *fibonacci_api.FibonacciRequest) (*fibonacci_api.FibonacciResponse, error) {
	x := req.GetX()
	y := req.GetY()
	if x > y {
		/*todo: error*/
		return nil, errors.New("wrong indexes: x must be less than y")
	}
	if y > 2000 {
		return nil, errors.New("y must be less than 2000")
	}
	return &fibonacci_api.FibonacciResponse{Slice: createFibonacciInterval(x, y)}, nil
}
