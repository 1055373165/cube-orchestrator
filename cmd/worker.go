/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cube/worker"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Worker command to operate a Cube worker node.",
	Long: `worker is a entity for execute task:

1. Worker receive manager scheduler request with task and running task on docker container.
2. Worker responds to the manager's request abount task states
3. Worker collects container cpu, disk and memory detail which serve to schedule task.
`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		name, _ := cmd.Flags().GetString("name")
		dbType, _ := cmd.Flags().GetString("dbType")

		log.Println("starting worker.")
		w := worker.New(name, dbType)
		api := worker.Api{Address: host, Port: port, Worker: w}
		go w.RunTasks()
		go w.CollectStats()
		go w.UpdateTasks()

		log.Printf("Starting worker API on http://%s:%d\n", host, port)
		api.Start()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.Flags().StringP("host", "H", "0.0.0.0", "Hostname or ip address")
	workerCmd.Flags().IntP("port", "p", 5000, "worker listening port")
	workerCmd.Flags().StringP("name", "n", fmt.Sprintf("worker-%s", uuid.New().String()), "worker unique identifier")
	workerCmd.Flags().StringP("dbtype", "d", "memory", "type of datastore to use tasks (\"memory\" or \"persistent\")")
}
