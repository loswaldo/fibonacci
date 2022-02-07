FROM golang:1.17

WORKDIR /go/src

ADD ./pkg ./pkg

ADD ./servermain ./servermain

ADD go.mod .

ADD go.sum .

RUN go build -o ./bin/grpc-fibonacci ./servermain

ENTRYPOINT ["/go/src/bin/grpc-fibonacci"]

EXPOSE 8080