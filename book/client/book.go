package main

import (
	"context"
	"log"

	pb "github.com/sikderg/go-grpc-example/book/proto"
)

func getBook(c pb.BookServiceClient) {
	log.Printf("Getting book")

	res, err := c.Book(context.Background(), &pb.BookRequest{
		Id: 123,
	})

	if err != nil {
		log.Fatalf("Failed to get book: %v", err)
	}

	log.Printf("Book ID: %v", res.Id)
	log.Printf("Book Title:%v", res.Title)

}
