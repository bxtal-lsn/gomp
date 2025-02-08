package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec <instance> <command>",
	Short: "Run a command inside a Multipass instance",
	Long: `
	Executes a command inside a Multipass instance.

	Examples:
	  # Run 'ls -l' inside an instance:
	  gomp exec my-instance ls -l
	
	  # Run a shell command:
	  gomp exec my-instance bash -c "echo Hello World"
	
	This is equivalent to:
	  multipass exec my-instance ls -l

`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		instance := args[0]
		command := args[1:]

		fmt.Printf("Running command in instance: %s\n", instance)

		execCmd := exec.Command("multipass", append([]string{"exec", instance, "--"}, command...)...)
		execCmd.Stdin = os.Stdin
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		err := execCmd.Run()
		if err != nil {
			fmt.Printf("Error running command in instance %s: %v\n", instance, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
