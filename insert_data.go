package main

import (
	"fmt"
	"log"
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

	// Create table if not exists
	err = session.Query(`
		CREATE TABLE IF NOT EXISTS records (
			id UUID PRIMARY KEY,
			date timestamp,
			data text
		)
	`).Exec()
	if err != nil {
		log.Fatal(err)
	}

	// Insert data for December 2024
	startDate := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	uniqueID := gocql.TimeUUID() // Use UUID for unique IDs

	for currentDate := startDate; currentDate.Before(endDate); currentDate = currentDate.Add(24 * time.Hour) {
		data := fmt.Sprintf("Data for %s", currentDate.Format("2006-01-02"))
		err := session.Query(`INSERT INTO records (id, date, data) VALUES (?, ?, ?)`, uniqueID, currentDate, data).Exec()
		if err != nil {
			log.Fatal(err)
		}
		uniqueID = gocql.TimeUUID() // Increment the unique ID
	}

	fmt.Println("Data insertion complete")
}
