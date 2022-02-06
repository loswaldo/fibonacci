package main

import (
	"github.com/loswaldo/fibonacci/pkg/fibonacci"
	"github.com/loswaldo/fibonacci/pkg/fibonacci_api"
	"google.golang.org/grpc"
	"log"
	"net"
)

//import "google.golang.org/grpc"

func main() {
	s := grpc.NewServer()
	srv := &fibonacci.GRPCServer{}
	fibonacci_api.RegisterFibonacciSliceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
