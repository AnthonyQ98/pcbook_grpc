package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"pcbook/pb"
	"pcbook/sample"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createLaptop(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	fmt.Printf("created laptop with id: %s\n", res.Id)
}

func searchLaptop(laptopClient pb.LaptopServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		createLaptop(laptopClient)
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		laptop := res.GetLaptop()
		log.Print("- found: ", laptop.GetId())
		log.Printf("  - price: %v", laptop.GetPriceUsd())
		log.Printf("  - cpu: %v", laptop.GetCpu())
		log.Printf("  - ram: %v", laptop.GetRam())
		log.Printf("  - name: %v", laptop.GetName())
		log.Printf("  - brand: %v", laptop.GetBrand())
		log.Printf("  - screen: %v", laptop.GetScreen())
		log.Printf("  - keyboard: %v", laptop.GetKeyboard())
		log.Printf("  - gpus: %v", laptop.GetGpus())
		log.Printf("  - storages: %v", laptop.GetStorages())
		log.Print("  - price: ", laptop.GetPriceUsd(), "usd")
	}
}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()

	fmt.Printf("start client with server address: %s\n", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	searchLaptop(laptopClient)

}
