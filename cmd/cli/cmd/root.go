package cmd

import (
	"log"
	"os"

	pb "github.com/s0lar/antib-bruteforce/gen/antibruteforce"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetGrpcClient - общий метод для инициализации grpc клиента.
func GetGrpcClient() pb.CheckerClient {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(cfg.App.ServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	grpcClient := pb.NewCheckerClient(conn)

	return grpcClient
}
