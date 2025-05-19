package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=127.0.0.1 user=postgres password=postgres dbname=postgres_work_with_db port=5435 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connect to database")
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping database")
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database!")
}
