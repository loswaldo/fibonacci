FROM golang:1.17

WORKDIR /go/src
#RUN mkdir -p "app/servermain"

#ADD ./pkg ./pkg
#ADD ./servermain ./servermain
#ADD go.mod ./go/src
#ADD go.sum ./go/src
#ENV GO111MODULE=auto
#RUN cd go/src && go mod download && go mod verify
#COPY ./pkg .
#COPY ./servermain .
#COPY go.mod .
#COPY go.sum .
COPY . .

#RUN cd go/src && ls


RUN go build -o ./bin/grpc-fibonacci ./servermain

CMD ["/go/src/bin/grpc-fibonacci"]

#RUN rm -rf clientmain && rm -rf proto
#COPY servermain/servermain.go ./servermain/

#COPY ./servermain/servermain.go  ./server/

#RUN mkdir -p "usr/local/go/src/github.com/loswaldo/fibonacci/pkg/fibonacci_api" \
#&&  mkdir -p "usr/local/go/src/github.com/loswaldo/fibonacci/pkg/fibonacci"
#
#COPY ./pkg/fibonacci/grpcserver.go ./src/github.com/loswaldo/fibonacci/pkg/fibonacci
#COPY ./pkg/fibonacci_api/fibonacci.pb.go ./src/github.com/loswaldo/fibonacci/pkg/fibonacci_api
#
#RUN GO111MODULE=on go mod download
#
#RUN go get google.golang.org/grpc
#
#ENTRYPOINT go build server/servermain.go
#
EXPOSE 5300