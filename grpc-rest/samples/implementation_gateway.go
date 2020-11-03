package samples

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	simple.RegisterHelloServer(srv, &server{})
	reflection.Register(srv)
	go srv.Serve(lis) // HL

	mux := runtime.NewServeMux()                                                                  // HL
	opts := []grpc.DialOption{grpc.WithInsecure()}                                                // HL
	simple.RegisterHelloHandlerFromEndpoint(context.Background(), mux, lis.Addr().String(), opts) // HL
	http.ListenAndServe(":8080", mux)                                                             // HL
}
