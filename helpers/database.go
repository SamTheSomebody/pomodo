package helpers

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"pomodo/internal/database"
)

func GetDBQueries() *database.Queries {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatalf("Error getting DB Queries: %v", err)
	}
	return database.New(db)
}
