package main

import (
	"context"
	"log"

	pb "github.com/sikderg/go-grpc-example/book/proto"
)

func (s *Server) Book(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {

	log.Printf("Received: %v", req)
	return &pb.BookResponse{
		//response BookResponse id and title from request BookRequest
		Id:    req.Id,
		Title: "Sample Book Title",
	}, nil

}
