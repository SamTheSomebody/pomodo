/*
Copyright © 2025 Sam Muller gamedevsam@pm.me
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"pomodo/helpers"
)

var (
	summary      string
	dueAt        string
	timeEstimate string
	priority     int
	enthusiasm   int
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a task",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.AddTask(task)
		if err != nil {
			log.Fatalf("SQL instertion error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&task.Summary, "summary", "s", "", "Add a summary for the task")
	addCmd.Flags().StringVarP(&task.DueAt, "due", "d", "", "Add a due date and/or time for the task")
	addCmd.Flags().StringVarP(&task.TimeEstimate, "estimate", "t", "", "Add a time estimate for the task")
	addCmd.Flags().StringVarP(&task.Priority, "priority", "p", "", "Add a priority level for the task (1-10)")
	addCmd.Flags().StringVarP(&task.Enthusiasm, "enthusiasm", "e", "", "Add an enthusiasm level for the task (1-10)")
}
