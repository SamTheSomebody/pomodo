package ui

import (
	"fmt"
	"time"

	"pomodo/internal/database"
)

func PrintTask(task database.Task) {
	output := task.Name + " - "
	output += fmt.Sprintf("Due: %v ", task.DueAt.Format(time.Stamp))
	output += fmt.Sprintf("Created: %v\n", task.CreatedAt.Format(time.Stamp))
	timeSpent := secondsToDuration(task.TimeSpentSeconds)
	output += fmt.Sprintf("Estimated time: %v", secondsToDuration(task.TimeEstimateSeconds))
	output += fmt.Sprintf(" (%v spent)\n", timeSpent)
	fmt.Println(output)
}

func secondsToDuration(seconds int64) string {
	return (time.Duration(seconds) * time.Second).String()
}
