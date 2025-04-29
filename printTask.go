package main

import (
	"database/sql"
	"fmt"
	"strings"

	"pomodo/internal/database"
)

func printTask(task database.Task) {
	fmt.Printf("%v - Due: %v, Created At: %v\nEstimated Time: %v (%v spent)\n%v\n",
		strings.ToUpper(task.Name), task.DueAt, task.CreatedAt,
		secondsToDuration(task.TimeEstimateSeconds), secondsToDuration(task.TimeSpentSeconds),
		task.Description)
}

func secondsToDuration(seconds sql.NullInt64) string {
	if !seconds.Valid {
		return "nil"
	}

	remainder := seconds.Int64
	output := ""
	if remainder < 60 {
		return fmt.Sprintf("%v seconds", remainder)
	}

	for remainder > 0 {
		mod := remainder % 60
		segment := ""
		if mod < 10 {
			segment = "0"
		}
		segment += fmt.Sprintf("%v", mod)
		output = segment + ":" + output
		remainder /= 60
	}

	return output
}
