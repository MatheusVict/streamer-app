package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	connectionString := "host=postgres port=5432 user=postgres dbname=streamkeys password=123 sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	return db, nil
}
