package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"pomodo/internal/database"
)

type state struct {
	DB       *database.Queries
	CFG      *Config
	Settings *Settings
	CLI      *cli
}

type Config struct {
	IsDebugMode bool
}

func InitializeState(isDebugMode bool) *state {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	if isDebugMode {
		fmt.Println("[DEBUG] Creating state")
	}
	return &state{
		DB: dbQueries,
		CFG: &Config{
			IsDebugMode: true,
		},
		Settings: NewSettings(), // TODO Read settings from file
		CLI:      &cli{},
	}
}
