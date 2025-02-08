package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount <host-folder> <instance>:<mount-path>",
	Short: "Mount a local folder inside a Multipass instance",
	Long: `
	Mounts a folder from the host system into a Multipass instance.

	Example:
	  gomp mount /home/user/my-data my-instance:/mnt/shared
	
	This is equivalent to:
	  multipass mount /home/user/my-data my-instance:/mnt/shared

`, Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()
		hostFolder := args[0]
		mountPath := args[1]

		fmt.Printf("Mounting %s to %s\n", hostFolder, mountPath)

		mountCmd := exec.Command("multipass", "mount", hostFolder, mountPath)
		mountCmd.Stdout = os.Stdout
		mountCmd.Stderr = os.Stderr

		err := mountCmd.Run()
		if err != nil {
			fmt.Printf("Error mounting folder: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}
