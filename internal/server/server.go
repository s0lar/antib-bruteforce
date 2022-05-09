package server

import (
	"context"
	"github.com/s0lar/antib-bruteforce/gen/antibruteforce"
)

type GRPCServer struct {
	//logger *zap.Logger
}

func (s *GRPCServer) Check(context.Context, *antibruteforce.CheckRequest) (*antibruteforce.CheckResponse, error) {
	return &antibruteforce.CheckResponse{Result: true}, nil
}
