package helpers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	"pomodo/internal/database"
)

type RawTask struct {
	ID           string
	Name         string
	Summary      string
	DueAt        string
	TimeEstimate string
	TimeSpent    string
	Enthusiasm   string
	Priority     string
}

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

func Raw(task database.Task) RawTask {
	id := ""
	switch task.ID.(type) {
	case string:
		id = task.ID.(string)
	case uuid.UUID:
		id = task.ID.(uuid.UUID).String()
	}
	rawTask := RawTask{
		ID:           id,
		Name:         task.Name,
		Summary:      task.Summary,
		DueAt:        defaultRawTime(task.DueAt),
		TimeEstimate: defaultRawDuration(task.TimeEstimateSeconds),
		TimeSpent:    defaultRawDuration(task.TimeSpentSeconds),
		Priority:     defaultRawRange(task.Priority),
		Enthusiasm:   defaultRawRange(task.Enthusiasm),
	}
	return rawTask
}

func defaultRawTime(t time.Time) string {
	x := t.Format(time.RFC822)
	if x == "01 Jan 01 00:00 UTC" {
		return ""
	}
	return x
}

func defaultRawDuration(d int64) string {
	if d == 0 {
		return ""
	}
	return time.Duration(d).String()
}

func defaultRawRange(i int64) string {
	if i == 0 {
		return ""
	}
	return strconv.Itoa(int(i))
}

func (t RawTask) Validate() (database.Task, error) {
	var err error
	task := database.Task{
		ID:      t.ID,
		Name:    t.Name,
		Summary: t.Summary,
	}

	if len(t.DueAt) != 0 {
		task.DueAt, err = time.Parse(time.RFC822, t.DueAt)
		if err != nil {
			return database.Task{}, fmt.Errorf("due at parse error: %v", err)
		}
	}

	if len(t.TimeEstimate) != 0 {
		d, err := time.ParseDuration(t.TimeEstimate)
		if err != nil {
			return database.Task{}, fmt.Errorf("time estimate parse error: %v", err)
		}
		task.TimeEstimateSeconds = int64(d.Seconds())
	}

	task.Priority = ValidateRange(t.Priority)
	task.Enthusiasm = ValidateRange(t.Enthusiasm)

	return task, nil
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

func AddTask(task RawTask) error {
	t, err := task.Validate()
	if err != nil {
		return fmt.Errorf("error validating task: %v", err)
	}
	db := GetDBQueries()
	params := database.CreateTaskParams{
		ID:                  uuid.New(),
		Name:                t.Name,
		Summary:             t.Summary,
		DueAt:               t.DueAt,
		TimeEstimateSeconds: t.TimeEstimateSeconds,
		Priority:            t.Priority,
		Enthusiasm:          t.Enthusiasm,
	}
	_, err = db.CreateTask(context.TODO(), params)
	return err
}

func EditTask(task RawTask) error {
	t, err := task.Validate()
	if err != nil {
		return fmt.Errorf("error validating task: %v", err)
	}
	db := GetDBQueries()
	params := database.SetTaskParams{
		Name:                t.Name,
		Summary:             t.Summary,
		DueAt:               t.DueAt,
		TimeEstimateSeconds: t.TimeEstimateSeconds,
		Priority:            t.Priority,
		Enthusiasm:          t.Enthusiasm,
	}
	_, err = db.SetTask(context.TODO(), params)
	return err
}
