package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias <alias-name>=<command>",
	Short: "Create an alias for a command inside a Multipass instance",
	Long: `
	Defines a shortcut (alias) for a command inside a Multipass instance.

	Aliases allow users to quickly execute frequently used commands.
	
	Examples:
	  # Create an alias for updating packages:
	  gomp alias update="sudo apt update && sudo apt upgrade -y"
	
	  # Create an alias for starting a web server:
	  gomp alias start-web="systemctl start nginx"
	
	To remove an alias, use:
	  unalias <alias-name>
`, Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		alias := args[0]

		fmt.Printf("Creating alias: %s\n", alias)

		aliasCmd := exec.Command("multipass", "alias", alias)
		aliasCmd.Stdout = os.Stdout
		aliasCmd.Stderr = os.Stderr

		err := aliasCmd.Run()
		if err != nil {
			fmt.Printf("Error creating alias: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
