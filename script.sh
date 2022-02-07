#!/bin/bash

docker build -t fibonacci-grpc-server .

docker run --rm -p 8080:8080 fibonacci-grpc-server