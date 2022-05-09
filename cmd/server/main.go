package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	srv := &antibruteforce.


	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error on Listen %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error on Serve gRPC %v", err)
	}
}
