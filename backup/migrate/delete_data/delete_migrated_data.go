package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

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

	// Open the backup CSV file
	file, err := os.Open(`D:\All Document iServeU\Test\data-migration-task\backup\backup.csv`)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Iterate over the rows and delete records from Cassandra
	for i, row := range records {
		// Skip the header row
		if i == 0 {
			continue
		}

		// The ID is in the first column of the CSV
		id, err := gocql.ParseUUID(row[0])
		if err != nil {
			log.Printf("Failed to parse UUID '%s': %v", row[0], err)
			continue
		}

		// Delete the record with the given ID
		err = session.Query(`DELETE FROM records WHERE id = ?`, id).Exec()
		if err != nil {
			log.Printf("Failed to delete record with ID '%s': %v", id, err)
		}
	}

	fmt.Println("Records present in the backup file have been deleted from Cassandra.")
}
