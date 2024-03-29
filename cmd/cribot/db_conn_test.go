package main_test 

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestDbConnection(t *testing.T) {
    if err := godotenv.Load("../../.env"); err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
    os.Getenv("HOST"),
    os.Getenv("PORT"),
    os.Getenv("USER"),
    os.Getenv("DBNAME"),
    os.Getenv("PASSWORD"),
)

fmt.Printf("%s\n", connStr)

db, err := sql.Open("postgres", connStr)
if err != nil {
    t.Fatal(err)
}

defer db.Close()

err = db.Ping(); if err != nil {
    t.Fatal(err)
}

if err != nil {
    t.Fatal(err)
}

t.Log("Database Connected!")
}
