package cmd

import (
	"context"
	"log"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/spf13/cobra"
)

// blacklistAddCmd represents the blacklistAdd command.
var blacklistAddCmd = &cobra.Command{
	Use:   "blacklistAdd",
	Short: "Add ip net into black list",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient := GetGrpcClient()
		res, err := grpcClient.AddBlacklist(context.Background(), &pb.NetListRequest{
			Net: cmd.Flag("net").Value.String(),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Result: %v", res.GetOk())
	},
}

func init() {
	rootCmd.AddCommand(blacklistAddCmd)
	blacklistAddCmd.PersistentFlags().String("net", "", "--net=172.17.0.0/16")
}
