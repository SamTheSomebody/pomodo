package helpers

import (
	"context"
	"log"

	"github.com/google/uuid"

	db "pomodo/database"
	"pomodo/internal/database"
)

func GetTask(context context.Context, value string) database.Task {
	db := db.GetDBQueries()
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
