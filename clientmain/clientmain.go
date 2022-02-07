package main

import (
	"context"
	"github.com/loswaldo/fibonacci/pkg/fibonacci_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := fibonacci_api.NewFibonacciSliceClient(conn)
	res, err := client.Fibonacci(context.Background(), &fibonacci_api.FibonacciRequest{X: 0, Y: 5})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetSlice())

}
