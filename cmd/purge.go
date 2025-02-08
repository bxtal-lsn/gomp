package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "🗑️ Purge unused Multipass images",
	Long: `🗑️ This command removes all unused Multipass images.
It does NOT delete any running or stopped instances.

Usage:
  gomp purge
`,
	Run: func(cmd *cobra.Command, args []string) {
		runMultipassPurge()
	},
}

func runMultipassPurge() {
	fmt.Println("🗑️ Running `multipass purge` to remove unused images...")

	cmd := exec.Command("multipass", "purge")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("❌ Error purging images:", err)
		return
	}

	fmt.Println("✅ Unused Multipass images have been purged.")
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
