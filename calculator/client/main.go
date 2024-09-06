package main

import (
	"log"

	pb "github.com/sikderg/go-grpc-example/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:8082"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection Failed %v", err)
	}

	log.Printf("Connection Status %v", conn.GetState())
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	getSum(c)
	getSub(c)
}
