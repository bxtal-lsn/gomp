package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:   "unmount <instance>:<mount-path>",
	Short: "Unmount a folder from a Multipass instance",
	Long: `
	Unmounts a folder that was previously mounted into a Multipass instance.

	Example:
	  gomp unmount my-instance:/mnt/shared

`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		mountPath := args[0]

		fmt.Printf("Unmounting %s\n", mountPath)

		unmountCmd := exec.Command("multipass", "umount", mountPath)
		unmountCmd.Stdout = os.Stdout
		unmountCmd.Stderr = os.Stderr

		err := unmountCmd.Run()
		if err != nil {
			fmt.Printf("Error unmounting folder: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)
}
