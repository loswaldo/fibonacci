package main

import (
	"context"
	"encoding/json"
	"github.com/loswaldo/fibonacci/pkg/fibonacci_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type message struct {
	Seq []uint64
}

func makeFibonacciHandler(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		log.Fatal(err)
	}
	sequence, err := makeFibonacciRequest(int64(x), int64(y))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	} else {
		m := message{Seq: sequence}
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = html.Execute(w, html)
	if err != nil {
		log.Fatal(err)
	}
}

func makeFibonacciRequest(x, y int64) ([]uint64, error) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := fibonacci_api.NewFibonacciSliceClient(conn)
	res, err := client.Fibonacci(context.Background(), &fibonacci_api.FibonacciRequest{X: x, Y: y})
	if err != nil {
		return nil, err
	}
	return res.GetSlice(), nil
}

func main() {

	http.HandleFunc("/fibonacci", fibonacciHandler)
	http.HandleFunc("/newfibonacci", makeFibonacciHandler)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		log.Fatal(err)
	}

}
