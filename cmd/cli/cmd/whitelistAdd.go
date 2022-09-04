package cmd

import (
	"context"
	"log"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/spf13/cobra"
)

// whitelistAddCmd represents the whitelistAdd command.
var whitelistAddCmd = &cobra.Command{
	Use:   "whitelistAdd",
	Short: "Add ip net into white list",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient := GetGrpcClient()
		res, err := grpcClient.AddWhitelist(context.Background(), &pb.NetListRequest{
			Net: cmd.Flag("net").Value.String(),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Result: %v", res.GetOk())
	},
}

func init() {
	rootCmd.AddCommand(whitelistAddCmd)
	whitelistAddCmd.PersistentFlags().String("net", "", "--net=172.17.0.0/16")
}
