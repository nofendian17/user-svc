package cmd

import (
	"auth-svc/src/cmd/http"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	restCmd := &cobra.Command{
		Use:   "rest",
		Short: "rest server",
		Long:  `Command to serve rest server `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Serving rest server")
			http.Start()
		},
	}

	rootCmd.AddCommand(restCmd)
}
