package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// insertPubKeyCmd represents the insert-pub-key command
var insertPubKeyCmd = &cobra.Command{
	Use:   "insert-pub-key <instance-name> [more-instance-names...]",
	Short: "Insert your SSH public key into one or more Multipass instances",
	Long: `Copies your local SSH public key (~/.ssh/id_rsa.pub) into the 
~/.ssh/authorized_keys file of one or more Multipass instances.

Example:
  gomp insert-pub-key my-instance
  gomp insert-pub-key instance1 instance2 instance3
`,
	Args: cobra.MinimumNArgs(1),
	Run:  insertPubKey,
}

func insertPubKey(cmd *cobra.Command, args []string) {
	pubKeyPath := os.ExpandEnv("$HOME/.ssh/id_rsa.pub")

	// Check if SSH public key exists
	if _, err := os.Stat(pubKeyPath); os.IsNotExist(err) {
		fmt.Println("💥 No SSH key found at ~/.ssh/id_rsa.pub. Please generate one with `ssh-keygen`.")
		return
	}

	// Read SSH public key
	pubKey, err := os.ReadFile(pubKeyPath)
	if err != nil {
		fmt.Println("💥 Error reading SSH key:", err)
		return
	}
	sshKey := strings.TrimSpace(string(pubKey))

	// Loop through all provided instances and insert the SSH key
	for _, instance := range args {
		fmt.Printf("🐧 Inserting SSH key into %s...\n", instance)

		// Ensure .ssh directory exists in the instance
		exec.Command("multipass", "exec", instance, "--", "mkdir", "-p", "$HOME/.ssh").Run()
		exec.Command("multipass", "exec", instance, "--", "chmod", "700", "$HOME/.ssh").Run()

		// Append the SSH key to authorized_keys
		cmd := exec.Command("multipass", "exec", instance, "--", "bash", "-c",
			fmt.Sprintf("echo '%s' >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys", sshKey))

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("💥 Failed to insert SSH key into %s: %v\n", instance, err)
		} else {
			fmt.Printf("🐧 SSH key successfully added to %s\n", instance)
		}
	}
}

func init() {
	rootCmd.AddCommand(insertPubKeyCmd)
}
