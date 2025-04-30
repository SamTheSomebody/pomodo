/*
Copyright © 2025 Sam Muller gamedevsam@pm.me
*/
package cmd

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	db "pomodo/database"
	"pomodo/internal/database"
	"pomodo/ui"
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
	Use:   "add",
	Short: "Add a task",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		params := database.CreateTaskParams{
			ID:   uuid.New(),
			Name: strings.TrimSpace(args[0]),
		}

		if len(summary) != 0 {
			params.Description = sql.NullString{
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

		if priority != 0 {
			if priority > 10 {
				priority = 10
			} else if priority < 1 {
				priority = 1
			}
			params.Priority = sql.NullInt64{
				Valid: true,
				Int64: int64(priority),
			}
		}

		if enthusiasm != 0 {
			if enthusiasm > 10 {
				enthusiasm = 10
			} else if enthusiasm < 1 {
				enthusiasm = 1
			}
			params.Enthusiasm = sql.NullInt64{
				Valid: true,
				Int64: int64(enthusiasm),
			}
		}

		db := db.GetDBQueries()
		task, err := db.CreateTask(cmd.Context(), params)
		if err != nil {
			log.Fatalf("SQL instertion error: %v", err)
		}
		ui.PrintTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&summary, "summary", "s", "", "Add a summary for the task")
	addCmd.Flags().StringVarP(&dueAt, "due", "d", "", "Add a due date and/or time for the task")
	addCmd.Flags().StringVarP(&timeEstimate, "estimate", "t", "", "Add a time estimate for the task")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 5, "Add a priority level (1-10) for the task")
	addCmd.Flags().IntVarP(&enthusiasm, "enthusiasm", "e", 5, "Add an enthusiasm level (1-10) for the task")
}
