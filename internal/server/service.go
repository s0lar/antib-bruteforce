package server

import (
	"context"
	"errors"
	"github.com/s0lar/antib-bruteforce/internal/netlist"
	"log"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/s0lar/antib-bruteforce/internal/bucket"
)

var ValidateEmptyIPError = errors.New("validate error: empty IP")
var ValidateWrongIPError = errors.New("validate error: wrong IP")
var ValidateEmptyLoginError = errors.New("validate error: empty Login")
var ValidateEmptyPasswordError = errors.New("validate error: empty Password")

type Server struct {
	pb.UnimplementedCheckerServer
	bucketLogin    *bucket.Bucket
	bucketPassword *bucket.Bucket
	bucketIP       *bucket.Bucket
	listWhite      *netlist.NetList
	listBlack      *netlist.NetList
	logger         log.Logger
}

func NewServer(bucketLogin, bucketPassword, bucketIP *bucket.Bucket, listWhite, listBlack *netlist.NetList) *Server {
	return &Server{
		bucketLogin:    bucketLogin,
		bucketPassword: bucketPassword,
		bucketIP:       bucketIP,
		listWhite:      listWhite,
		listBlack:      listBlack,
	}
}

//	Check - method check
func (s *Server) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	log.Printf("Check: Request (%s)\n", req)

	//	TODO. Validate
	if req.GetIp() == "" {
		log.Printf("Ok: false. %v", ValidateEmptyIPError)
		return &pb.CheckResponse{Ok: false}, ValidateEmptyIPError
	}

	if req.GetLogin() == "" {
		log.Printf("Ok: false. %v", ValidateEmptyLoginError)
		return &pb.CheckResponse{Ok: false}, ValidateEmptyLoginError
	}

	if req.GetPassword() == "" {
		log.Printf("Ok: false. %v", ValidateEmptyPasswordError)
		return &pb.CheckResponse{Ok: false}, ValidateEmptyPasswordError
	}

	//	TODO. Check IP

	if req.GetIp() != "" && s.listWhite.Find(req.GetIp()) {

	}

	if !s.bucketLogin.Allow(req.GetLogin()) {
		log.Printf("Ok: false. Bucket %s, Value %s\n", "login", req.GetLogin())
		return &pb.CheckResponse{Ok: false}, nil
	}
	if !s.bucketPassword.Allow(req.GetPassword()) {
		log.Printf("Ok: false. Bucket %s, Value %s\n", "password", req.GetPassword())
		return &pb.CheckResponse{Ok: false}, nil
	}
	if !s.bucketIP.Allow(req.GetIp()) {
		log.Printf("Ok: false. Bucket %s, Value %s\n", "ip", req.GetIp())
		return &pb.CheckResponse{Ok: false}, nil
	}

	log.Printf("Ok: true\n")
	return &pb.CheckResponse{Ok: true}, nil
}

func (s *Server) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.ResetResponse, error) {
	log.Printf("Reset: Request (%s)\n", req)

	if req.GetLogin() == "" {
		log.Printf("Ok: false. Empty login")
		return &pb.ResetResponse{Ok: true}, nil
	}

	if req.GetPassword() == "" {
		log.Printf("Ok: false. Empty password")
		return &pb.ResetResponse{Ok: true}, nil
	}

	s.bucketLogin.Reset(req.GetLogin())
	s.bucketPassword.Reset(req.GetPassword())

	log.Printf("Ok: true\n")
	return &pb.ResetResponse{Ok: true}, nil
}

func (s *Server) mustEmbedUnimplementedCheckerServer() {
	panic("implement me")
}
