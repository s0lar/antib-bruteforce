package server

import (
	"context"
	"errors"
	"log"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/s0lar/antib-bruteforce/internal/bucket"
	"github.com/s0lar/antib-bruteforce/internal/netlist"
)

var (
	ErrorValidateEmptyIP       = errors.New("validate error: empty IP")
	ErrorValidateEmptyLogin    = errors.New("validate error: empty Login")
	ErrorValidateEmptyPassword = errors.New("validate error: empty Password")
	ErrorValidateEmptyNet      = errors.New("validate error: empty net")
)

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

	if req.GetIp() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyIP)
		return &pb.CheckResponse{Ok: false}, ErrorValidateEmptyIP
	}

	if req.GetLogin() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyLogin)
		return &pb.CheckResponse{Ok: false}, ErrorValidateEmptyLogin
	}

	if req.GetPassword() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyPassword)
		return &pb.CheckResponse{Ok: false}, ErrorValidateEmptyPassword
	}

	//	Check IP in WhiteList. If found then return Ok:true
	if s.listWhite.Find(req.GetIp()) {
		log.Printf("Ok: true. IP found in white list %s\n", req.GetIp())
		return &pb.CheckResponse{Ok: true}, nil
	}

	//	Check IP in BlackList. If found then return Ok:false
	if s.listBlack.Find(req.GetIp()) {
		log.Printf("Ok: false. IP found in black list %s\n", req.GetIp())
		return &pb.CheckResponse{Ok: false}, nil
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
		log.Printf("Ok: false. %v", ErrorValidateEmptyLogin)
		return &pb.ResetResponse{Ok: false}, ErrorValidateEmptyLogin
	}

	if req.GetPassword() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyPassword)
		return &pb.ResetResponse{Ok: false}, ErrorValidateEmptyPassword
	}

	s.bucketLogin.Reset(req.GetLogin())
	s.bucketPassword.Reset(req.GetPassword())

	log.Printf("Ok: true\n")
	return &pb.ResetResponse{Ok: true}, nil
}

func (s *Server) AddWhitelist(ctx context.Context, req *pb.NetListRequest) (*pb.NetListResponse, error) {
	if req.GetNet() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyNet)
		return &pb.NetListResponse{Ok: false}, ErrorValidateEmptyNet
	}

	if err := s.listWhite.Add(req.GetNet()); err != nil {
		log.Printf("Ok: false. %v", err)
		return &pb.NetListResponse{Ok: false}, err
	}

	return &pb.NetListResponse{Ok: true}, nil
}

func (s *Server) RemoveWhitelist(ctx context.Context, req *pb.NetListRequest) (*pb.NetListResponse, error) {
	if req.GetNet() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyNet)
		return &pb.NetListResponse{Ok: false}, ErrorValidateEmptyNet
	}

	if err := s.listWhite.Remove(req.GetNet()); err != nil {
		log.Printf("Ok: false. %v", err)
		return &pb.NetListResponse{Ok: false}, err
	}

	return &pb.NetListResponse{Ok: true}, nil
}

func (s *Server) AddBlacklist(ctx context.Context, req *pb.NetListRequest) (*pb.NetListResponse, error) {
	if req.GetNet() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyNet)
		return &pb.NetListResponse{Ok: false}, ErrorValidateEmptyNet
	}

	if err := s.listBlack.Add(req.GetNet()); err != nil {
		log.Printf("Ok: false. %v", err)
		return &pb.NetListResponse{Ok: false}, err
	}

	return &pb.NetListResponse{Ok: true}, nil
}

func (s *Server) RemoveBlacklist(ctx context.Context, req *pb.NetListRequest) (*pb.NetListResponse, error) {
	if req.GetNet() == "" {
		log.Printf("Ok: false. %v", ErrorValidateEmptyNet)
		return &pb.NetListResponse{Ok: false}, ErrorValidateEmptyNet
	}

	if err := s.listBlack.Remove(req.GetNet()); err != nil {
		log.Printf("Ok: false. %v", err)
		return &pb.NetListResponse{Ok: false}, err
	}

	return &pb.NetListResponse{Ok: true}, nil
}

func (s *Server) mustEmbedUnimplementedCheckerServer() {
	panic("implement me")
}
