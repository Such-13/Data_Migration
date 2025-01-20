package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	// Connect to Cassandra
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "test_keyspace"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Open CSV file for writing
	file, err := os.Create("backup.csv")
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	headers := []string{"ID", "Date", "Data"}
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Failed to write headers to CSV: %v", err)
	}

	// Query to fetch data (last 10 days of December 2024)
	startDate := time.Date(2024, 12, 22, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	query := `SELECT id, date, data FROM records WHERE date >= ? AND date <= ? ALLOW FILTERING`
	iter := session.Query(query, startDate, endDate).Iter()

	// Variables to store the results
	var id gocql.UUID
	var date time.Time
	var data string

	for iter.Scan(&id, &date, &data) {
		// Write each row to the CSV file
		row := []string{
			id.String(),
			date.Format("2006-01-02 15:04:05"),
			data,
		}
		if err := writer.Write(row); err != nil {
			log.Fatalf("Failed to write row to CSV: %v", err)
		}
	}

	// Check for errors during iteration
	if err := iter.Close(); err != nil {
		log.Fatalf("Failed to close iterator: %v", err)
	}

	fmt.Println("Data backup completed and saved to backup.csv")
}
