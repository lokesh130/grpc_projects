package main

import (
	"fmt"
	"log"
	"net"
	"flag"
	"context"
	pb "calculator_app/calculator_def"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 40001, "Port to run calculator service at")
)

type server struct { 
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, numPr *pb.NumPair) (*pb.Result, error) { 
	return &pb.Result{Result: numPr.Num1 + numPr.Num2}, nil
}

func (s *server) Sub(ctx context.Context, numPr *pb.NumPair) (*pb.Result, error) { 
	return &pb.Result{Result: numPr.Num1 - numPr.Num2}, nil
}

func (s *server) Mul(ctx context.Context, numPr *pb.NumPair) (*pb.Result, error) { 
	return &pb.Result{Result: numPr.Num1 * numPr.Num2}, nil
}

func (s *server) Div(ctx context.Context, numPr *pb.NumPair) (*pb.Result, error) { 
	return &pb.Result{Result: numPr.Num1 / float64(numPr.Num2)}, nil
}

func main() { 
	flag.Parse()
	log.Printf("Starting Calculator Server .....")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("Not able to create listner on locahost at port: %d. Error: %v", *port, err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCalculatorServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil { 
		log.Fatal("Unable to run grpc-server on listner. Failed with err: %v", err)
	}
}