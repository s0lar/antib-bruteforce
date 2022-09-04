package cmd

import (
	"context"
	"log"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/spf13/cobra"
)

// blacklistRemoveCmd represents the blacklistRemove command.
var blacklistRemoveCmd = &cobra.Command{
	Use:   "blacklistRemove",
	Short: "Remove ip net from black list",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient := GetGrpcClient()
		res, err := grpcClient.RemoveBlacklist(context.Background(), &pb.NetListRequest{
			Net: cmd.Flag("net").Value.String(),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Result: %v", res.GetOk())
	},
}

func init() {
	rootCmd.AddCommand(blacklistRemoveCmd)
	blacklistRemoveCmd.PersistentFlags().String("net", "", "--net=172.17.0.0/16")
}
