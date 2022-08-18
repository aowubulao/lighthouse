#!/bin/sh
cd ..
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./target/lighthouse-server-amd64 main.go
