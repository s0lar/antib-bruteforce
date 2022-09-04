package cmd

import (
	"context"
	"log"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/spf13/cobra"
)

// bucketResetCmd represents the bucketReset command.
var bucketResetCmd = &cobra.Command{
	Use:   "bucketReset",
	Short: "Reset buckets",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient := GetGrpcClient()
		res, err := grpcClient.Reset(context.Background(), &pb.ResetRequest{
			Login:    cmd.Flag("login").Value.String(),
			Password: cmd.Flag("password").Value.String(),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Result: %v", res.GetOk())
	},
}

func init() {
	rootCmd.AddCommand(bucketResetCmd)

	bucketResetCmd.PersistentFlags().String("login", "", "--login=DoeJoe")
	bucketResetCmd.PersistentFlags().String("password", "", "--password=secret")
}
