package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all running and stopped Multipass instances",
	Long: `
	Displays a list of all Multipass instances, including their:
	  - Name
	  - Status (Running / Stopped)
	  - IP Address
	  - Memory & CPU allocation
	
	Example:
	  multipass-cli list

`, Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		out, err := exec.Command("multipass", "list").Output()
		if err != nil {
			fmt.Println("Error listing instances:", err)
			return
		}
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
