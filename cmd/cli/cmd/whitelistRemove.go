/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"log"

	"github.com/spf13/cobra"
)

// whitelistRemoveCmd represents the whitelistRemove command
var whitelistRemoveCmd = &cobra.Command{
	Use:   "whitelistRemove",
	Short: "Remove ip net from white list",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient := GetGrpcClient()
		res, err := grpcClient.RemoveWhitelist(context.Background(), &pb.NetListRequest{
			Net: cmd.Flag("net").Value.String(),
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Result: %v", res.GetOk())
	},
}

func init() {
	rootCmd.AddCommand(whitelistRemoveCmd)
	whitelistRemoveCmd.PersistentFlags().String("net", "", "--net=172.17.0.0/16")
}
