package main

import (
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}
}

func main() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Server(lis)
}
