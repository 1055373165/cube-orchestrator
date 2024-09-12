/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a new task",
	Long: `cube run command.
	
The run command starts a new task.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager, _ := cmd.Flags().GetString("manager")
		filename, _ := cmd.Flags().GetString("filename")
		fullFilePath, err := filepath.Abs(filename)
		if err != nil {
			log.Fatal(err)
		}
		if !fileExist(fullFilePath) {
			log.Fatalf("File %s doesn't exist.", filename)
		}

		log.Printf("Using manager: %v\n", manager)
		log.Printf("Using file: %v\n", fullFilePath)

		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("Unable to read file: %v", filename)
		}
		log.Printf("Data: %v\n", string(data))

		url := fmt.Sprintf("http://%s/tasks", manager)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Panic(err)
		}

		if resp.StatusCode != http.StatusCreated {
			log.Printf("Error sending request: %v", resp.StatusCode)
		}

		defer resp.Body.Close()
		log.Println("Successfully sent task request to manager")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("manager", "m", "localhost:5556", "Manager to task to")
	runCmd.Flags().StringP("filename", "f", "task.json", "Task specification")
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, fs.ErrNotExist)
}
