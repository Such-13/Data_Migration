# Data Migration Project

This project is designed to manage data migration tasks between Cassandra and PostgreSQL using Go. It includes scripts for data backup, migration, and cleanup, with the help of Dockerized environments.

## Project Structure

```
D:.
├── docker-compose.yml          # Defines the Docker services for Cassandra and PostgreSQL
├── Dockerfile                  # Dockerfile for Go environment
├── go.mod                      # Go module file
├── go.sum                      # Go dependencies
├── insert_data.go              # Script to insert test data into Cassandra
│
├───backup
│   ├── backup.csv              # Backup file generated from Cassandra
│   ├── backup_data.go          # Script to back up data from Cassandra
│   └───migrate
│       ├── migrate_to_postgres.go  # Script to migrate data to PostgreSQL
│       └───delete_data
│           └── delete_migrated_data.go  # Script to delete migrated data from Cassandra
```

## Prerequisites

- **Go**: Make sure Go is installed (version 1.20 or later recommended).
- **Docker**: Ensure Docker and Docker Compose are installed.
- **Databases**: Cassandra and PostgreSQL images are defined in the `docker-compose.yml` file.

## Setup and Usage

### 1. Clone the Repository

```bash
git clone <repository-url>
cd data-migration
```

### 2. Start Docker Containers

Start the database containers using Docker Compose:

```bash
docker-compose up -d
```

### 3. Install Dependencies

Install Go dependencies defined in `go.mod`:

```bash
go mod tidy
```

### 4. Insert Sample Data

Run the script to insert sample data into Cassandra:

```bash
go run insert_data.go
```

### 5. Backup Data from Cassandra

Export data from Cassandra into a CSV file:

```bash
go run backup/backup_data.go
```

### 6. Migrate Data to PostgreSQL

Import the backup data into PostgreSQL:

```bash
go run backup/migrate/migrate_to_postgres.go
```

### 7. Delete Migrated Data from Cassandra

Clean up the data already migrated to PostgreSQL:

```bash
go run backup/migrate/delete_data/delete_migrated_data.go
```

## Configuration

- **PostgreSQL**: The default database name is `postgres` with password `1234`. Update the credentials if necessary in the scripts.
- **Cassandra**: Default configuration is defined in the `docker-compose.yml` file.

## Notes

- Ensure the Docker containers are running before executing the Go scripts.
- Logs and errors are printed to the console for debugging purposes.

## Cleanup

Stop and remove the Docker containers:

```bash
docker-compose down
```

## Contributing

Feel free to contribute by submitting issues or pull requests.

## License

This project is licensed under the me. 

