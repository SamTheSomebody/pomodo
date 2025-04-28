package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	"pomodo/internal/database"
)

/*
type CreateTaskParams struct {
	ID                  interface{}
	Name                string
	Description         sql.NullString
	DueAt               sql.NullTime
	TimeEstimateSeconds sql.NullInt64
	Priority            sql.NullInt64
	Enthusiasm          sql.NullInt64
} */

func addTaskHandler(s *state, c command) error {
	params := database.CreateTaskParams{
		ID:   uuid.New(),
		Name: c.arguments["default"],
	}

	if value, ok := c.arguments["-d"]; ok {
		time, err := time.Parse(time.Stamp, value)
		if err != nil {
			log.Fatal("Unable to parse due at: " + value)
		}
		params.DueAt = sql.NullTime{
			Time: time, Valid: true,
		}
	}

	if value, ok := c.arguments["-t"]; ok {
		duration, err := time.ParseDuration(value)
		if err != nil {
			log.Fatal("Unable to parse time estimate: " + value)
		}
		params.TimeEstimateSeconds = sql.NullInt64{
			Int64: int64(duration.Seconds()), Valid: true,
		}
	}

	if value, ok := c.arguments["-p"]; ok {
		amount, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			log.Fatal("Unable to parse priority: " + value)
		}
		if amount > 10 {
			amount = 10
		} else if amount < 0 {
			amount = 0
		}
		params.Priority = sql.NullInt64{
			Int64: amount, Valid: true,
		}
	}

	if value, ok := c.arguments["-e"]; ok {
		amount, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			log.Fatal("Unable to parse enthusiasm: " + value)
		}
		if amount > 10 {
			amount = 10
		} else if amount < 0 {
			amount = 0
		}
		params.Priority = sql.NullInt64{
			Int64: amount, Valid: true,
		}
	}

	if value, ok := c.arguments["-s"]; ok {
		params.Description = sql.NullString{
			String: value, Valid: true,
		}
	}

	s.DB.CreateTask(context.Background(), params)
	return nil
}
