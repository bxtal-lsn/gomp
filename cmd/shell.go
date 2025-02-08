package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell <instance>",
	Short: "Open an interactive shell inside a Multipass instance",
	Long: `
	Opens a shell session inside a running Multipass instance.

	Example:
	  multipass-cli shell my-instance

	This is equivalent to:
	  multipass shell my-instance
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		instance := args[0]

		fmt.Printf("Opening shell in instance: %s\n", instance)

		shellCmd := exec.Command("multipass", "shell", instance)
		shellCmd.Stdin = os.Stdin
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr

		err := shellCmd.Run()
		if err != nil {
			fmt.Printf("Error opening shell in instance %s: %v\n", instance, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
