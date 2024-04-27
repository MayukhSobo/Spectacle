package cmd

import (
	"spectacle/internal/app"
	"spectacle/log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spectacle",
	Short: "Spectacle is an ETCD explorer for your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger.Info("Starting Spectacle")
		app.Start()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
