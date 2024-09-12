/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cube/task"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/docker/go-units"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status comamnd to list tasks.",
	Long: `cube status command.
	
The status command allow a user to get the status of tasks from
the Cube manager`,
	Run: func(cmd *cobra.Command, args []string) {
		manager, _ := cmd.Flags().GetString("manager")

		url := fmt.Sprintf("http://%s/tasks", manager)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()

		var tasks []*task.Task
		err = json.NewDecoder(resp.Body).Decode(&tasks)
		if err != nil {
			log.Fatal(err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "ID\tNAME\tCREATED\tSTATE\tCONTAINERNAME\tIMAGE\t")
		for _, task := range tasks {
			var start string
			if task.StartTime.IsZero() {
				start = fmt.Sprintf("%s ago", units.HumanDuration(time.Now().UTC().Sub(time.Now().UTC()))) // Human Duration 返回一个人类可读的持续时间近似值（"About a minute", "4 hours ago",）。
			} else {
				start = fmt.Sprintf("%s ago", units.HumanDuration(time.Now().UTC().Sub(task.StartTime)))
			}

			// TODO: there is a bug here, state for stopped jobs is showing as Running
			state := task.State.StateStringSlice()[task.State]
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t\n", task.ID, task.Name, start, state, task.Name, task.Image)
		}
		// Flush should be called after the last call to Write to ensure that
		// any data buffered in the Writer is written to output.
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("manager", "m", "localhost:5556", "manager to talk to")
}
