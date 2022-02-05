package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"pkg/fibonacci_api/pkg/fibonacci_api/pkg/fibonacci_api"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
