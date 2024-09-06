package main

import (
	"context"
	"log"

	pb "github.com/sikderg/go-grpc-example/calculator/proto"
)

func (s *Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Request Reveived %v", req)
	return &pb.AddResponse{
		Sum: req.Num1 + req.Num2,
	}, nil
}

func (s *Server) Sub(ctx context.Context, req *pb.SubRequest) (*pb.SubResponse, error) {
	log.Printf("Request for Subtraction %v", req)

	return &pb.SubResponse{
		Sub: req.Num1 - req.Num2,
	}, nil
}
