package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var deleteAll bool

var deleteCmd = &cobra.Command{
	Use:   "delete <instance> | --all",
	Short: "Delete a Multipass instance (or all instances)",
	Long: `
	Deletes one or more Multipass instances. If an instance is deleted but not purged, 
	it can still be recovered using the 'recover' command.
	
	Examples:
	  # Delete a single instance:
	  gomp delete my-instance
	
	  # Delete multiple instances:
	  gomp delete instance1 instance2
	
	  # Delete ALL instances (WARNING: Irreversible!):
	  gomp delete --all

`, Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()

		if deleteAll {
			fmt.Println("Deleting all instances...")
			exec.Command("multipass", "delete", "--all", "--purge").Run()
			return
		}

		for _, name := range args {
			fmt.Printf("Deleting instance: %s\n", name)
			exec.Command("multipass", "delete", name, "--purge").Run()
		}
	},
}

func init() {
	deleteCmd.Flags().BoolVar(&deleteAll, "all", false, "Delete all instances")
	rootCmd.AddCommand(deleteCmd)
}
