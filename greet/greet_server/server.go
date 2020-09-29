package main

import (
	"fmt"
	"log"
	"net"

	"greet/greetpb"
	"C:\Users\LENOVO\go\src\greet\greetpb"

	"github.com/simplesteph/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	//k = server.server1{}
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
