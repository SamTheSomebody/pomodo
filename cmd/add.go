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
	"pomodo/helpers"
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
	Use:   "add [name]",
	Short: "Add a task",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		params := database.CreateTaskParams{
			ID:   uuid.New(),
			Name: strings.TrimSpace(args[0]),
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
	addCmd.Flags().IntVarP(&priority, "priority", "p", 5, "Add a priority level for the task (1-10)")
	addCmd.Flags().IntVarP(&enthusiasm, "enthusiasm", "e", 5, "Add an enthusiasm level for the task (1-10)")
}
