package main

import (
	"log"
	"net"
	"time"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/s0lar/antib-bruteforce/internal/bucket"
	"github.com/s0lar/antib-bruteforce/internal/netlist"
	"github.com/s0lar/antib-bruteforce/internal/server"
	"google.golang.org/grpc"
)

func main() {
	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	interval := 1000 * time.Second
	ttl := interval * 2

	listWhite := netlist.NewNetList()
	listBlack := netlist.NewNetList()

	if err := listWhite.Load([]string{"192.168.10.0/24"}); err != nil {
		log.Fatal(err)
	}
	if err := listBlack.Load([]string{"172.168.10.0/24"}); err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	srv := server.NewServer(
		bucket.NewBucket("login", 10, interval, ttl),
		bucket.NewBucket("password", 1000, interval, ttl),
		bucket.NewBucket("ip", 1000, interval, ttl),
		listWhite,
		listBlack,
	)
	pb.RegisterCheckerServer(grpcServer, srv)

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := grpcServer.Serve(lsn); err != nil {
		log.Fatal(err)
	}

	// grpcServer := server.NewServer()
	// srv := &antibruteforce.
	//
	//
	// lis, err := net.Listen("tcp", ":9000")
	// if err != nil {
	//	log.Fatalf("Error on Listen %v", err)
	//}
	//
	// if err := grpcServer.Serve(lis); err != nil {
	//	log.Fatalf("Error on Serve gRPC %v", err)
	//}
}
