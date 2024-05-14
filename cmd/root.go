package cmd

import (
	"spectacle/internal/app/home"
	"spectacle/logger"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spectacle",
	Short: "Spectacle is an ETCD explorer for your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log.Info("Starting Spectacle")
		home.Start()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
