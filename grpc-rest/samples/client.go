package client

import (
	"context"
	"fmt"
	"log"

	grpc "google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewGreeterClient(cc)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Boston"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Message)
}
