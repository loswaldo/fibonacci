FROM golang:1.12

RUN mkdir "server"

COPY go.mod ./server/
COPY go.sum ./server/

COPY ./servermain/servermain.go  ./server/

RUN mkdir -p "usr/local/go/src/github.com/loswaldo/fibonacci/pkg/fibonacci_api" \
&&  mkdir -p "usr/local/go/src/github.com/loswaldo/fibonacci/pkg/fibonacci"

COPY ./pkg/fibonacci/grpcserver.go ./src/github.com/loswaldo/fibonacci/pkg/fibonacci
COPY ./pkg/fibonacci_api/fibonacci.pb.go ./src/github.com/loswaldo/fibonacci/pkg/fibonacci_api

RUN GO111MODULE=on go mod download

RUN go get google.golang.org/grpc

ENTRYPOINT go build server/servermain.go

EXPOSE 5300