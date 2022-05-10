package server

import (
	"context"
	"github.com/s0lar/antib-bruteforce/internal/netlist"
	"testing"
	"time"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/s0lar/antib-bruteforce/internal/bucket"
	"github.com/stretchr/testify/require"
)

func TestServer_Check(t *testing.T) {
	var ctx context.Context
	limit := 10
	interval := 1 * time.Minute
	ttl := interval * 2
	testsCount := 11

	// srv := NewServer(
	//	server.NewBucket("login", limit, interval, ttl),
	//	server.NewBucket("password", limit, interval, ttl),
	//	server.NewBucket("ip", limit, interval, ttl),
	//)

	tests := []struct {
		name    string
		server  *Server
		okTrue  int
		okFalse int
	}{
		{
			"Check login bucket",
			NewServer(
				bucket.NewBucket("login", limit, interval, ttl),
				bucket.NewBucket("password", testsCount+1, interval, ttl),
				bucket.NewBucket("ip", testsCount+1, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			limit,
			testsCount - limit,
		},
		{
			"Check password bucket",
			NewServer(
				bucket.NewBucket("login", testsCount+1, interval, ttl),
				bucket.NewBucket("password", limit, interval, ttl),
				bucket.NewBucket("ip", testsCount+1, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			limit,
			testsCount - limit,
		},
		{
			"Check ip bucket",
			NewServer(
				bucket.NewBucket("login", testsCount+1, interval, ttl),
				bucket.NewBucket("password", testsCount+1, interval, ttl),
				bucket.NewBucket("ip", limit, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			limit,
			testsCount - limit,
		},
		{
			"Zero limit",
			NewServer(
				bucket.NewBucket("login", 0, interval, ttl),
				bucket.NewBucket("password", 0, interval, ttl),
				bucket.NewBucket("ip", 0, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			0,
			testsCount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			okTrue := 0
			okFalse := 0
			for i := 0; i < testsCount; i++ {
				res, _ := tt.server.Check(ctx, &pb.CheckRequest{
					Login:    "login",
					Password: "password",
					Ip:       "192.168.0.1",
				})

				if res.GetOk() {
					okTrue++
				} else {
					okFalse++
				}
			}

			require.Equal(t, tt.okTrue, okTrue)
			require.Equal(t, tt.okFalse, okFalse)
		})
	}

	//
	// type fields struct {
	//	UnimplementedCheckerServer antibruteforce.UnimplementedCheckerServer
	//	bucketLogin                *server.Bucket
	//	bucketPassword             *server.Bucket
	//	bucketIP                   *server.Bucket
	//	logger                     log.Logger
	//}
	// type args struct {
	//	ctx context.Context
	//	req *pb.CheckRequest
	//}
	// tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	want    *pb.CheckResponse
	//	wantErr bool
	// }{
	//	// TODO: Add test cases.
	//}
	// for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		s := &Server{
	//			UnimplementedCheckerServer: tt.fields.UnimplementedCheckerServer,
	//			bucketLogin:                tt.fields.bucketLogin,
	//			bucketPassword:             tt.fields.bucketPassword,
	//			bucketIP:                   tt.fields.bucketIP,
	//			logger:                     tt.fields.logger,
	//		}
	//		got, err := s.Check(tt.args.ctx, tt.args.req)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("Check() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
