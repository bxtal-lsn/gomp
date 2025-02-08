package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Create and configure new Multipass instances",
	Long: `
	Launch one or more Multipass Ubuntu instances.

	You can specify:
	  - Instance names (space-separated)
	  - Memory, storage, and CPU allocation per instance
	  - Cloud-init YAML for provisioning
	
	Examples:
	  # Launch a single instance with default settings:
	  gomp launch
	  
	  # Launch multiple instances with shared config:
	  gomp launch
	  > Enter instance names: test1 test2 test3
	  > Use the same config for all instances? yes
	  > Memory: 4GB
	  > Storage: 10GB
	  > CPUs: 2
	
	  # Launch instances with different configs:
	  gomp launch
	  > Enter instance names: test1 test2
	  > Use the same config for all instances? no
	  > Configuring test1...
	  > Memory: 2GB
	  > Storage: 8GB
	  > CPUs: 1
	  > Configuring test2...
	  > Memory: 4GB
	  > Storage: 10GB
	  > CPUs: 2

`,
	Run: func(cmd *cobra.Command, args []string) {
		checkMultipass()

		// Ask for server names
		fmt.Print("🐧 Enter instance names (space-separated): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		names := strings.Fields(scanner.Text())

		if len(names) == 0 {
			fmt.Println("No instance names provided. Exiting.")
			return
		}

		var instancesConfig []InstanceConfig

		// If only one instance, just prompt for its config
		if len(names) == 1 {
			fmt.Printf("Configuring instance: %s\n", names[0])
			memory, storage, cpus := getResourceConfig()
			cloudInitPath := getCloudInitConfig()
			instancesConfig = append(instancesConfig, InstanceConfig{
				Name:      names[0],
				Memory:    memory,
				Storage:   storage,
				CPUs:      cpus,
				CloudInit: cloudInitPath,
			})
		} else {
			// If multiple instances, ask if batch config should be used
			fmt.Print("Use the same config for all instances? (yes/no): ")
			scanner.Scan()
			useBatch := strings.ToLower(scanner.Text()) == "yes"

			var memory, storage, cpus, cloudInitPath string
			if useBatch {
				memory, storage, cpus = getResourceConfig()
				cloudInitPath = getCloudInitConfig()

				for _, name := range names {
					instancesConfig = append(instancesConfig, InstanceConfig{
						Name:      name,
						Memory:    memory,
						Storage:   storage,
						CPUs:      cpus,
						CloudInit: cloudInitPath,
					})
				}
			} else {
				// If not using batch, prompt for each instance separately
				for _, name := range names {
					fmt.Printf("\nConfiguring instance: %s\n", name)
					mem, stor, cpu := getResourceConfig()
					cloudInitPath := getCloudInitConfig()
					instancesConfig = append(instancesConfig, InstanceConfig{
						Name:      name,
						Memory:    mem,
						Storage:   stor,
						CPUs:      cpu,
						CloudInit: cloudInitPath,
					})
				}
			}
		}

		// Launch all instances after gathering configs
		for _, config := range instancesConfig {
			cmdArgs := []string{"launch", "-n", config.Name, "--memory", config.Memory, "--disk", config.Storage, "--cpus", config.CPUs}
			if config.CloudInit != "" {
				cmdArgs = append(cmdArgs, "--cloud-init", config.CloudInit)
			}

			fmt.Printf("Launching instance: %s with %s RAM, %s disk, %s CPUs, Cloud-Init: %s\n",
				config.Name, config.Memory, config.Storage, config.CPUs, config.CloudInit)

			// Run the actual Multipass command and check for errors
			launchCmd := exec.Command("multipass", cmdArgs...)
			launchCmd.Stdout = os.Stdout
			launchCmd.Stderr = os.Stderr
			err := launchCmd.Run()
			if err != nil {
				fmt.Printf("Error launching instance %s: %v\n", config.Name, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}

func checkMultipass() {
	_, err := exec.LookPath("multipass")
	if err != nil {
		fmt.Println("Multipass is not installed. Please install it first.")
		os.Exit(1)
	}
}

type InstanceConfig struct {
	Name      string
	Memory    string
	Storage   string
	CPUs      string
	CloudInit string
}

func getResourceConfig() (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter memory (default 2GB): ")
	scanner.Scan()
	memory := formatGB(scanner.Text(), "2GB")

	fmt.Print("Enter storage (default 10GB): ")
	scanner.Scan()
	storage := formatGB(scanner.Text(), "10GB")

	fmt.Print("Enter CPUs (default 2): ")
	scanner.Scan()
	cpus := scanner.Text()
	if cpus == "" {
		cpus = "2"
	}

	return memory, storage, cpus
}

// Allows each instance to have a different cloud-init file, or use the default.
func getCloudInitConfig() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Use cloud-init? (yes/no): ")
	scanner.Scan()
	useCloudInit := strings.ToLower(scanner.Text()) == "yes"

	if !useCloudInit {
		return ""
	}

	cloudInitPath := "./cloudinit.yml"
	fmt.Print("Enter cloud-init YAML file path (or press Enter for default ./cloudinit.yml): ")
	scanner.Scan()
	input := scanner.Text()
	if input != "" {
		cloudInitPath = input
	}

	// Validate cloud-init file
	if _, err := os.Stat(cloudInitPath); os.IsNotExist(err) {
		fmt.Println("Invalid file path. Skipping cloud-init for this instance.")
		return ""
	}

	return cloudInitPath
}

// formatGB ensures input has "GB" at the end. If empty, returns default value.
func formatGB(input, defaultValue string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	if !strings.HasSuffix(input, "GB") {
		return input + "GB"
	}
	return input
}
