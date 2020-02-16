FROM golang:latest

WORKDIR /go/src/github.com/julioshinoda/transfer-api

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/rest/server.go" --command=./server