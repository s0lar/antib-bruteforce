package netlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	cidr  string
	isErr bool
}{
	{
		"172.17.0.0/16",
		false,
	},
	{
		"172.17.0.0/16",
		false,
	},
	{
		"192.17.0.0/24",
		false,
	},
	{
		"192.17.0.0",
		true,
	},
	{
		"asd",
		true,
	},
	{
		"asd",
		true,
	},
}

func TestNetList_Add(t *testing.T) {
	list := NewNetList()

	for _, tt := range tests {
		t.Run(tt.cidr, func(t *testing.T) {
			err := list.Add(tt.cidr)
			if err != nil {
				require.True(t, tt.isErr)
			}
		})
	}
}

func TestNetList_Find(t *testing.T) {
	list := NewNetList()

	for _, tt := range tests {
		list.Add(tt.cidr)
	}

	tests := []struct {
		ip     string
		result bool
	}{
		{
			"172.17.0.1",
			true,
		},
		{
			"192.17.0.100",
			true,
		},
		{
			"192.17.0.200",
			true,
		},
		{
			"asd",
			false,
		},
		{
			"80.17.0.0",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			require.Equal(t, tt.result, list.Find(tt.ip))
		})
	}
}

func TestNetList_Load(t *testing.T) {
	tlist := NewNetList()

	err := tlist.Load([]string{"90.17.0.0/24", "100.17.0.0/24"})
	require.NoError(t, err)

	err = tlist.Load([]string{"asdasd"})
	require.Error(t, err)
}

func TestNetList_Remove(t *testing.T) {
	list := NewNetList()

	//	Add some CIDR's before removing
	for _, tt := range tests {
		t.Run(tt.cidr, func(t *testing.T) {
			list.Add(tt.cidr)
		})
	}

	for _, tt := range tests {
		t.Run(tt.cidr, func(t *testing.T) {
			err := list.Remove(tt.cidr)
			if err != nil {
				require.True(t, tt.isErr)
			}
		})
	}
}
