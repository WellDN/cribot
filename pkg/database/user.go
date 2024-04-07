package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/welldn/cribot/pkg/common"
)


func CreateTableUser(db *sql.DB) error {
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

func CreateDBUser(db *sql.DB, user common.DBUser) error { 
    _, err := db.Exec("INSERT INTO Users (ID, Name, Password) VALUES ($1, $2, $3)",
    user.ID, user.Name, user.Password)
    if err != nil {
        return err
    }
    fmt.Println("User created!!")
    return nil
}

func GetUserDByName(db *sql.DB, name string) (*common.DBUser, error) {
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

func GetUserDByID(db *sql.DB, id string) (*common.DBUser, error) {
    var user common.DBUser
    err := db.QueryRow("SELECT * FROM Users WHERE Name = $1", id).Scan(&user.ID, &user.Name, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("User not found")
        }
        return nil, err
    }
    return &user, nil
}
