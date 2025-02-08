package cmd

import (
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// getMultipassInstances fetches running instances from multipass list
func getMultipassInstances() []string {
	out, err := exec.Command("multipass", "list", "--format", "csv").Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(out), "\n")
	var instances []string
	for _, line := range lines[1:] { // Skip header line
		fields := strings.Split(line, ",")
		if len(fields) > 0 {
			instances = append(instances, strings.TrimSpace(fields[0]))
		}
	}
	return instances
}

// autoCompleteInstances provides dynamic instance name completion
func autoCompleteInstances(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return getMultipassInstances(), cobra.ShellCompDirectiveNoFileComp
}
