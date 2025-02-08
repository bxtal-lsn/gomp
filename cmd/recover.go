package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var recoverCmd = &cobra.Command{
	Use:   "recover <instance>",
	Short: "Recover a deleted (but not purged) Multipass instance",
	Long: `
	Restores an instance that was deleted but not purged.

	Example:
	  gomp recover my-instance
	
	To permanently delete an instance, use:
	  gomp delete my-instance --purge
	
`, Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		instance := args[0]

		fmt.Printf("Recovering instance: %s\n", instance)

		recoverCmd := exec.Command("multipass", "recover", instance)
		recoverCmd.Stdout = os.Stdout
		recoverCmd.Stderr = os.Stderr

		err := recoverCmd.Run()
		if err != nil {
			fmt.Printf("Error recovering instance %s: %v\n", instance, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(recoverCmd)
}
