/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// pruneCmd represents the prune command
var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "stop and remove container to avoid conflict",
	Long: `the quick way to stop and remove defineded container:

since the container to be started for each pending task is the same:
smy-test:v1, and the port exposed by the program running in the container
is 7777, the container needs to be cleand up when the program is restarted
to prevent port conflict.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopAndRemoveContainer()
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}

func stopAndRemoveContainer() {
	// 1. Execute "docker ps -aq" to get all container IDs
	cmd := exec.Command("docker", "ps", "-aq")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to list containers: %v", err)
	}

	// 2. Convert output to a slice of container IDs
	containerIDs := strings.Fields(string(output))
	if len(containerIDs) == 0 {
		fmt.Println("No containers are currently running.")
		return
	}

	// 3. Stop each container
	for _, containerID := range containerIDs {
		fmt.Printf("Stopping container %s...\n", containerID)
		stopCmd := exec.Command("docker", "stop", containerID)
		stopOutput, err := stopCmd.CombinedOutput() // output to stdout and stderr
		if err != nil {
			log.Printf("Failed to stop container %s: %v\n", containerID, err)
			continue
		}
		log.Printf("Container %s stopped: %s", containerID, stopOutput)
		removeCmd := exec.Command("docker", "rm", containerID)
		removeOutput, err := removeCmd.CombinedOutput()
		if err != nil {
			log.Printf("Failed to remove container %s: %v\n", containerID, err)
			continue
		}
		log.Printf("Container %s removed: %s", containerID, removeOutput)
	}
}
