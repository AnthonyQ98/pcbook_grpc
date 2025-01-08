package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"pcbook/pb"
	"pcbook/service"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the port number")
	flag.Parse()

	fmt.Printf("start server on port %d\n", *port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	laptopServer := service.NewLaptopServer(laptopStore, imageStore)
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
