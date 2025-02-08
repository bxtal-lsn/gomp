package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <instance>",
	Short: "Show detailed information about a Multipass instance",
	Long: `
	Displays detailed information about a Multipass instance, including:
	  - IP Address
	  - Disk & memory usage
	  - CPU allocation
	  - Running state
	
	Example:
	  gomp info my-instance

`, Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		instance := args[0]

		fmt.Printf("Fetching info for instance: %s\n", instance)

		infoCmd := exec.Command("multipass", "info", instance)
		infoCmd.Stdout = os.Stdout
		infoCmd.Stderr = os.Stderr

		err := infoCmd.Run()
		if err != nil {
			fmt.Printf("Error fetching info for instance %s: %v\n", instance, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
