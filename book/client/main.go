package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/sikderg/go-grpc-example/book/proto"
)

var addr string = "localhost:8081"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	defer conn.Close()
	c := pb.NewBookServiceClient(conn)
	getBook(c)

}
