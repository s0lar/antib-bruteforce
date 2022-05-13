package main

import (
	"database/sql"
	"log"
	"net"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/s0lar/antib-bruteforce/internal/bucket"
	"github.com/s0lar/antib-bruteforce/internal/netlist"
	"github.com/s0lar/antib-bruteforce/internal/server"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"google.golang.org/grpc"
)

func main() {
	//	TODO Config
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config loaded")

	//	TODO DB
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DB.Dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	log.Println("Database connected")
	log.Println(db)

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
		bucket.NewBucket(cfg.App.Login.Limit, cfg.App.Login.Interval, cfg.App.Login.TTL),          //	Login
		bucket.NewBucket(cfg.App.Password.Limit, cfg.App.Password.Interval, cfg.App.Password.TTL), //	Password
		bucket.NewBucket(cfg.App.IP.Limit, cfg.App.IP.Interval, cfg.App.IP.TTL),                   //	IP
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
