package main

import (
	"database/sql"
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

func InitializeState() *state {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	return &state{
		DB: dbQueries,
		CFG: &Config{
			IsDebugMode: true,
		},
		Settings: NewSettings(), // TODO Read settings from file
		CLI:      &cli{},
	}
}
