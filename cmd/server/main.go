package main

import (
	"log"
	"net"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/s0lar/antib-bruteforce/internal/bucket"
	"github.com/s0lar/antib-bruteforce/internal/netlist"
	"github.com/s0lar/antib-bruteforce/internal/server"
	"google.golang.org/grpc"
)

func main() {
	//	TODO Config
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	//	TODO DB
	//	TODO CLI
	//	TODO Docker && Make

	lsn, err := net.Listen("tcp", cfg.App.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}

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
		bucket.NewBucket(cfg.App.Login.limit, cfg.App.Login.interval, cfg.App.Login.ttl),          //	Login
		bucket.NewBucket(cfg.App.Password.limit, cfg.App.Password.interval, cfg.App.Password.ttl), //	Password
		bucket.NewBucket(cfg.App.IP.limit, cfg.App.IP.interval, cfg.App.IP.ttl),                   //	IP
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
