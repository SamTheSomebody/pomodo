/*
Copyright © 2025 Sam Muller gamedevsam@pm.me
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"pomodo/helpers"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a task",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		fetchedTask := helpers.GetTask(cmd.Context(), args[0])
		taskValidated, err := task.Validate()
		if err != nil {
			log.Fatal(err)
		}

		if task.Name != "" {
			fetchedTask.Name = taskValidated.Name
		}
		if task.Summary != "" {
			fetchedTask.Summary = taskValidated.Summary
		}
		if task.DueAt != "" {
			fetchedTask.Summary = taskValidated.Summary
		}
		if task.TimeEstimate != "" {
			fetchedTask.TimeEstimateSeconds = taskValidated.TimeEstimateSeconds
		}
		if task.TimeSpent != "" {
			fetchedTask.TimeEstimateSeconds = taskValidated.TimeSpentSeconds
		}
		if task.Priority != "" {
			fetchedTask.Priority = taskValidated.Priority
		}
		if task.Enthusiasm != "" {
			fetchedTask.Enthusiasm = taskValidated.Enthusiasm
		}
		err = helpers.EditTask(task)
		if err != nil {
			log.Fatalf("SQL instertion error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&task.Name, "name", "n", "", "edit the name of the task")
	editCmd.Flags().StringVarP(&task.Summary, "summary", "s", "", "edit a summary for the task")
	editCmd.Flags().StringVarP(&task.DueAt, "due", "d", "", "edit a due date and/or time for the task")
	editCmd.Flags().StringVarP(&task.TimeEstimate, "estimate", "t", "", "edit a time estimate for the task")
	editCmd.Flags().StringVarP(&task.Priority, "priority", "p", "", "edit a priority level for the task (1-10)")
	editCmd.Flags().StringVarP(&task.Enthusiasm, "enthusiasm", "e", "", "edit an enthusiasm level for the task (1-10)")
}
