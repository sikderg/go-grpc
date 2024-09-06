package main

import (
	"log"
	"net"

	pb "github.com/sikderg/go-grpc-example/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:8082"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Unable to Listen Port %v", err)
	}

	log.Printf("Server Listening at %v ", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %v", err)
	}
}
