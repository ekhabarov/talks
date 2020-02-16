package main

import (
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	grpc "google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	go s.Server(lis)

	handler := runtime.NewServeMux()
	pb.RegisterGreeterHandlerFromEndpoint(ctx, handler, lis.Addr().String())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
