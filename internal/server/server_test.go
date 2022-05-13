package server

import (
	"context"
	"testing"
	"time"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/s0lar/antib-bruteforce/internal/bucket"
	"github.com/s0lar/antib-bruteforce/internal/netlist"
	"github.com/stretchr/testify/require"
)

var netListTests = []struct {
	net    string
	isOk   bool
	isErr  bool
	errMsg string
}{
	{
		"172.17.0.0/16",

		true,
		false,
		"",
	},
	{
		"0.0.0.0",

		false,
		true,
		"invalid CIDR address: 0.0.0.0",
	},
	{
		"password",

		false,
		true,
		"invalid CIDR address: password",
	},
	{
		"",
		false,
		true,
		"validate error: empty net",
	},
}

func TestServer_Check(t *testing.T) {
	var ctx context.Context
	limit := 10
	interval := 1 * time.Minute
	ttl := interval * 2
	testsCount := 11

	netList := netlist.NewNetList()
	netList.Add("192.168.0.0/16")

	tests := []struct {
		name    string
		server  *Server
		okTrue  int
		okFalse int
	}{
		{
			"Check login bucket",
			NewServer(
				bucket.NewBucket(limit, interval, ttl),
				bucket.NewBucket(testsCount+1, interval, ttl),
				bucket.NewBucket(testsCount+1, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			limit,
			testsCount - limit,
		},
		{
			"Check password bucket",
			NewServer(
				bucket.NewBucket(testsCount+1, interval, ttl),
				bucket.NewBucket(limit, interval, ttl),
				bucket.NewBucket(testsCount+1, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			limit,
			testsCount - limit,
		},
		{
			"Check ip bucket",
			NewServer(
				bucket.NewBucket(testsCount+1, interval, ttl),
				bucket.NewBucket(testsCount+1, interval, ttl),
				bucket.NewBucket(limit, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			limit,
			testsCount - limit,
		},
		{
			"Zero limit",
			NewServer(
				bucket.NewBucket(0, interval, ttl),
				bucket.NewBucket(0, interval, ttl),
				bucket.NewBucket(0, interval, ttl),
				netlist.NewNetList(),
				netlist.NewNetList(),
			),
			0,
			testsCount,
		},
		{
			"White list",
			NewServer(
				bucket.NewBucket(0, interval, ttl),
				bucket.NewBucket(0, interval, ttl),
				bucket.NewBucket(0, interval, ttl),
				netList,
				netlist.NewNetList(),
			),
			testsCount,
			0,
		},
		{
			"Black list",
			NewServer(
				bucket.NewBucket(0, interval, ttl),
				bucket.NewBucket(0, interval, ttl),
				bucket.NewBucket(0, interval, ttl),
				netlist.NewNetList(),
				netList,
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
}

func TestServer_Reset(t *testing.T) {
	var ctx context.Context
	server := NewServer(
		bucket.NewBucket(10, 1*time.Minute, 2*time.Minute),
		bucket.NewBucket(10, 1*time.Minute, 2*time.Minute),
		bucket.NewBucket(10, 1*time.Minute, 2*time.Minute),
		netlist.NewNetList(),
		netlist.NewNetList(),
	)

	tests := []struct {
		name     string
		login    string
		password string
		isOk     bool
		isErr    bool
		err      error
	}{
		{
			"Reset login bucket",
			"login",
			"password",
			true,
			false,
			nil,
		},
		{
			"Reset login bucket",
			"",
			"password",
			false,
			true,
			ErrorValidateEmptyLogin,
		},
		{
			"Reset login bucket",
			"login",
			"",
			false,
			true,
			ErrorValidateEmptyPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := server.Reset(ctx, &pb.ResetRequest{
				Login:    tt.login,
				Password: tt.password,
			})

			require.Equal(t, tt.isOk, res.GetOk())

			if err != nil {
				require.True(t, tt.isErr)
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}

func TestServer_AddWhitelist(t *testing.T) {
	ctx, server := prepareCtxServer()

	for _, tt := range netListTests {
		t.Run(tt.net, func(t *testing.T) {
			res, err := server.AddWhitelist(ctx, &pb.NetListRequest{Net: tt.net})
			require.Equal(t, tt.isOk, res.GetOk())
			if err != nil {
				require.True(t, tt.isErr)
				require.Equal(t, tt.errMsg, err.Error())
			}
		})
	}
}

func TestServer_AddBlacklist(t *testing.T) {
	ctx, server := prepareCtxServer()

	for _, tt := range netListTests {
		t.Run(tt.net, func(t *testing.T) {
			res, err := server.AddBlacklist(ctx, &pb.NetListRequest{Net: tt.net})
			require.Equal(t, tt.isOk, res.GetOk())
			if err != nil {
				require.True(t, tt.isErr)
				require.Equal(t, tt.errMsg, err.Error())
			}
		})
	}
}

func TestServer_RemoveWhitelist(t *testing.T) {
	ctx, server := prepareCtxServer()

	for _, tt := range netListTests {
		t.Run(tt.net, func(t *testing.T) {
			res, err := server.RemoveWhitelist(ctx, &pb.NetListRequest{Net: tt.net})
			require.Equal(t, tt.isOk, res.GetOk())
			if err != nil {
				require.True(t, tt.isErr)
				require.Equal(t, tt.errMsg, err.Error())
			}
		})
	}
}

func TestServer_RemoveBlacklist(t *testing.T) {
	ctx, server := prepareCtxServer()

	for _, tt := range netListTests {
		t.Run(tt.net, func(t *testing.T) {
			res, err := server.RemoveBlacklist(ctx, &pb.NetListRequest{Net: tt.net})
			require.Equal(t, tt.isOk, res.GetOk())
			if err != nil {
				require.True(t, tt.isErr)
				require.Equal(t, tt.errMsg, err.Error())
			}
		})
	}
}

func prepareCtxServer() (context.Context, *Server) {
	var ctx context.Context
	server := NewServer(
		bucket.NewBucket(10, 1*time.Minute, 2*time.Minute),
		bucket.NewBucket(10, 1*time.Minute, 2*time.Minute),
		bucket.NewBucket(10, 1*time.Minute, 2*time.Minute),
		netlist.NewNetList(),
		netlist.NewNetList(),
	)

	return ctx, server
}
