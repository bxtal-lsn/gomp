package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start <instance>",
	Short: "Start a stopped Multipass instance",
	Long: `
	Starts one or more stopped Multipass instances.

	Example:
	  multipass-cli start my-instance

`, Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		for _, name := range args {
			fmt.Printf("Starting instance: %s\n", name)
			exec.Command("multipass", "start", name).Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
