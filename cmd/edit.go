/*
Copyright © 2025 Sam Muller gamedevsam@pm.me
*/
package cmd

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	db "pomodo/database"
	"pomodo/helpers"
	"pomodo/internal/database"
	"pomodo/ui"
)

var name string

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a task",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		task := helpers.GetTask(cmd.Context(), args[0])
		params := database.UpdateTaskParams{
			ID:                  task.ID,
			Name:                task.Name,
			Summary:             task.Summary,
			DueAt:               task.DueAt,
			TimeEstimateSeconds: task.TimeEstimateSeconds,
			Enthusiasm:          task.Enthusiasm,
			Priority:            task.Priority,
		}

		if len(name) != 0 {
			params.Name = strings.TrimSpace(name)
		}

		if len(summary) != 0 {
			params.Summary = sql.NullString{
				Valid:  true,
				String: strings.TrimSpace(summary),
			}
		}

		if len(dueAt) != 0 {
			t, err := time.Parse(time.RFC822, dueAt)
			if err != nil {
				log.Fatal(err)
			}
			params.DueAt = sql.NullTime{
				Valid: true,
				Time:  t,
			}
		}

		if len(timeEstimate) != 0 {
			d, err := time.ParseDuration(timeEstimate)
			if err != nil {
				log.Fatal(err)
			}
			params.TimeEstimateSeconds = sql.NullInt64{
				Valid: true,
				Int64: int64(d.Seconds()),
			}
		}

		params.Priority = helpers.ValidateRange(priority)
		params.Enthusiasm = helpers.ValidateRange(enthusiasm)

		db := db.GetDBQueries()
		task, err := db.UpdateTask(cmd.Context(), params)
		if err != nil {
			log.Fatalf("SQL instertion error: %v", err)
		}
		ui.PrintTask(task)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&name, "name", "n", "", "edit the name of the task")
	editCmd.Flags().StringVarP(&summary, "summary", "s", "", "edit a summary for the task")
	editCmd.Flags().StringVarP(&dueAt, "due", "d", "", "edit a due date and/or time for the task")
	editCmd.Flags().StringVarP(&timeEstimate, "estimate", "t", "", "edit a time estimate for the task")
	editCmd.Flags().IntVarP(&priority, "priority", "p", 5, "edit a priority level for the task (1-10)")
	editCmd.Flags().IntVarP(&enthusiasm, "enthusiasm", "e", 5, "edit an enthusiasm level for the task (1-10)")
}
