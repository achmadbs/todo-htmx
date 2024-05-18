package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var db *pgx.Conn

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get PostgreSQL connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connString := "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName

	// Connect to PostgreSQL
	var connectErr error
	db, connectErr = pgx.Connect(context.Background(), connString)
	if connectErr != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	db.Close(context.Background())
}
