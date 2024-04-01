package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/welldn/cribot/pkg/common"
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


func createTableUser(db *sql.DB) error {
    _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users (
        ID SERIAL PRIMARY KEY,
        Name VARCHAR (50) UNIQUE NOT NULL,
        Password VARCHAR (50) NOT NULL
    )`)
    if err != nil {
        return err
    }
    return nil
}

func createDBUser(db *sql.DB, user common.DBUser) error { 
    _, err := db.Exec("INSERT INTO Users (ID, Name, Password) VALUES ($1, $2, $3)",
    user.ID, user.Name, user.Password)
    if err != nil {
        return err
    }
    fmt.Println("User created!!")
    return nil
}

func getUserByName(db *sql.DB, name string) (*common.DBUser, error) {
    var user common.DBUser
    err := db.QueryRow("SELECT * FROM Users WHERE Name = $1", name).Scan(&user.ID, &user.Name, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("User not found")
        }
        return nil, err
    }
    return &user, nil
}
