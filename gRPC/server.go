package main

import (
	"log"
	"net"
	pb "./Intro/Intro_pb"
	"google.golang.org/grpc"
)

const (
	port = ":12345"
)

func main(){
	list, err := net.Listen("tcp", port)
	if err!=nil{
		log.Fatalln(" Cant listen on the port")
	}
	Server := grpc.NewServer()
	pb.
	log.Printf("Starting gRPC listener on port " + port)
	if err := Server.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}



