package main

import (
	"context"
	"log"

	pb "github.com/sikderg/go-grpc-example/calculator/proto"
)

func getSum(c pb.CalculatorServiceClient) {
	log.Printf("Geting Sum")

	res, err := c.Add(context.Background(), &pb.AddRequest{
		Num1: 12,
		Num2: 5,
	})

	if err != nil {
		log.Fatalf("Unable to Connect Calculator %v", err)
	}

	log.Printf("Result : %v", res.Sum)

}

func getSub(c pb.CalculatorServiceClient) {
	log.Printf("Getting Sub")

	res, err := c.Sub(context.Background(), &pb.SubRequest{
		Num1: 44,
		Num2: 11,
	})

	if err != nil {
		log.Fatalf("Unable to Connect Calculator %v", err)
	}
	log.Printf("Subtract : %v", res.Sub)
}
