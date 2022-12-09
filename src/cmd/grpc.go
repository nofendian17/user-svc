package cmd

import (
	"auth-svc/src/cmd/grpc"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rpcCmd := &cobra.Command{
		Use:   "grpc",
		Short: "grpc server",
		Long:  `Command to serve grpc server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Serving grpc server")
			grpc.Start()
		},
	}

	rootCmd.AddCommand(rpcCmd)
}
