package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQL connection string
	connStr := "postgres://postgres:1234@localhost:5433/postgres?sslmode=disable"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer db.Close()

	// Ensure the table exists in the `test_keyspace` schema
	_, err = db.Exec(`
		CREATE SCHEMA IF NOT EXISTS test_keyspace;
		CREATE TABLE IF NOT EXISTS test_keyspace.records (
			id UUID PRIMARY KEY,
			date TIMESTAMP,
			data TEXT
		)
	`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// Open the backup CSV file
	file, err := os.Open(`D:\All Document iServeU\Test\data-migration-task\backup\backup.csv`)
	if err != nil {
		log.Fatal("Failed to open backup file:", err)
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header row
	if err != nil {
		log.Fatal("Failed to read CSV header:", err)
	}

	// Insert data into PostgreSQL
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

		// Parse the date string from the CSV
		parsedDate, err := time.Parse("2006-01-02 15:04:05", record[1])
		if err != nil {
			log.Printf("Skipping record due to date parse error: %v", err)
			continue // Skip this record if date parsing fails
		}

		// Insert data into PostgreSQL
		_, err = db.Exec(`
			INSERT INTO test_keyspace.records (id, date, data) VALUES ($1, $2, $3)
		`, record[0], parsedDate, record[2])
		if err != nil {
			log.Printf("Failed to insert record (ID=%s): %v", record[0], err)
			continue // Skip and continue if there is an error for this record
		}

		log.Printf("Inserted record: ID=%s, Date=%v, Data=%s", record[0], parsedDate, record[2])
	}

	fmt.Println("Data migration to PostgreSQL complete!")
}
