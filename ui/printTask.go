package ui

import (
	"database/sql"
	"fmt"
	"time"

	"pomodo/internal/database"
)

func PrintTask(task database.Task) {
	output := task.Name + " - "
	if task.DueAt.Valid {
		output += fmt.Sprintf("Due: %v ", task.DueAt.Time.Format(time.Stamp))
	}
	output += fmt.Sprintf("Created: %v\n", task.CreatedAt.Format(time.Stamp))
	timeSpent := secondsToDuration(task.TimeSpentSeconds)
	if task.TimeEstimateSeconds.Valid {
		output += fmt.Sprintf("Estimated time: %v", secondsToDuration(task.TimeEstimateSeconds))
		if task.TimeSpentSeconds.Valid {
			output += fmt.Sprintf(" (%v spent)\n", timeSpent)
		} else {
			output += "\n"
		}
	} else if task.TimeSpentSeconds.Valid {
		output += fmt.Sprintf("Spent: %v\n", timeSpent)
	}
	fmt.Println(output)
}

func secondsToDuration(seconds sql.NullInt64) string {
	return (time.Duration(seconds.Int64) * time.Second).String()
}
