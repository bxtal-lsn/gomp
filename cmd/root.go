package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomp",
	Short: "A cli tool to manage Multipass Ubuntu instances",
	Long:  "A simple Go CLI tool using Cobra to create, list, start, stop, and delete Multipass instances.",
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
