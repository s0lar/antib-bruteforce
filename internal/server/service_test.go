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
		{
			"White list",
			NewServer(
				bucket.NewBucket("login", 0, interval, ttl),
				bucket.NewBucket("password", 0, interval, ttl),
				bucket.NewBucket("ip", 0, interval, ttl),
				netList,
				netlist.NewNetList(),
			),
			testsCount,
			0,
		},
		{
			"Black list",
			NewServer(
				bucket.NewBucket("login", 0, interval, ttl),
				bucket.NewBucket("password", 0, interval, ttl),
				bucket.NewBucket("ip", 0, interval, ttl),
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
	limit := 10
	interval := 1 * time.Minute
	ttl := interval * 2

	server := NewServer(
		bucket.NewBucket("login", limit, interval, ttl),
		bucket.NewBucket("password", limit, interval, ttl),
		bucket.NewBucket("ip", limit, interval, ttl),
		netlist.NewNetList(),
		netlist.NewNetList(),
	)

	tests := []struct {
		name     string
		login    string
		password string
		server   *Server
		isOk     bool
		isErr    bool
		err      error
	}{
		{
			"Reset login bucket",
			"login",
			"password",
			server,
			true,
			false,
			nil,
		},
		{
			"Reset login bucket",
			"",
			"password",
			server,
			false,
			true,
			ErrorValidateEmptyLogin,
		},
		{
			"Reset login bucket",
			"login",
			"",
			server,
			false,
			true,
			ErrorValidateEmptyPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.server.Reset(ctx, &pb.ResetRequest{
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
