package main

import (
	pb "calculator_app/calculator_def"
	"flag"

	// "fmt"
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:40001", "Addr of calcuatuor-grpc-server")
)

func main() { 
	flag.Parse()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(*addr, opts...)
	defer conn.Close()

	if err != nil { 
		log.Fatalf("Unalble to create conection with target: %v, got error: %v", *addr, err)
	}	

	client := pb.NewCalculatorClient(conn)

	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, _ := client.Add(ctx, &pb.NumPair{Num1: 10, Num2: 20})

	log.Printf("Got result of Add call. Result: %v", result)



	result1, err1 := client.Sub(ctx, &pb.NumPair{Num1: 10, Num2: 20})
	if err1 != nil { 
		log.Fatalf("Error in Sub. Error: %v", err1)
	}
	log.Printf("Got result of Sub call. Result: %v", result1)



	result2, err2 := client.Mul(ctx, &pb.NumPair{Num1: 10, Num2: 20})
	if err2 != nil { 
		log.Fatalf("Error in Sub. Error: %v", err2)
	}
	log.Printf("Got result of Mul call. Result: %v", result2)


	result3, err3 := client.Div(ctx, &pb.NumPair{Num1: 10, Num2: 20})
	if err3 != nil { 
		log.Fatalf("Error in Sub. Error: %v", err3)
	}
	log.Printf("Got result of Div call. Result: %v", result3)
}