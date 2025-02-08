package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop <instance>",
	Short: "Stop a running Multipass instance",
	Long: `
	Stops a currently running Multipass instance.

	Example:
	  gomp stop my-instance
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		for _, name := range args {
			fmt.Printf("Stopping instance: %s\n", name)
			exec.Command("multipass", "stop", name).Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
