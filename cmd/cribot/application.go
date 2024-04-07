package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {
    if err := godotenv.Load("../../.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // wcyd
        connEnv := "host=" + os.Getenv("HOST") +
        " port=" + os.Getenv("PORT") +
        " user=" + os.Getenv("USER") +
        " dbname=" + os.Getenv("DBNAME") +
        " password=" + os.Getenv("PASSWORD") +
        " sslmode=disable"

    db, err := sql.Open("postgres", connEnv)
    if err != nil {
        log.Fatal(err)
    }

    // Verify if the connection is alive, establishing a connection if necessary
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    return db, err
}

