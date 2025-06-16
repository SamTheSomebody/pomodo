package helpers

import (
	"context"
	"log"
	"pomodo/internal/database"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GetTask(context context.Context, value string) database.Task {
	db := GetDBQueries()
	var task database.Task
	var err error
	if id, e := uuid.Parse(value); e == nil {
		task, err = db.GetTaskByID(context, id)
	} else {
		task, err = db.GetTaskByName(context, value)
	}
	if err != nil {
		log.Fatalf("SQL error: %v", err)
	}
	return task
}

func ParseTime(t time.Time) string {
	x := t.Format(time.RFC822)
	if x == "01 Jan 01 00:00 UTC" {
		return ""
	}
	return x
}

func ParseDuration(d int64) string {
	if d == 0 {
		return ""
	}
	t := time.Duration(d) * time.Second
	return t.String()
}

func ValidateRange(value string) int64 {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	if i > 10 {
		return 10
	} else if i < 1 {
		return 1
	}
	return int64(i)
}

func ValidateDuration(s string) error {
	_, err := time.ParseDuration(s)
	return err
}

func ValidateTime(s string) error {
	_, err := time.Parse("YYYY/MM/DD", s)
	return err
}
