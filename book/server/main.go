package main

import (
	"log"
	"net"

	pb "github.com/sikderg/go-grpc-example/book/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:8081"

type Server struct {
	pb.BookServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)

	}

}
