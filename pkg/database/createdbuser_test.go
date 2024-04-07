package database_test

import (
	"database/sql"
	"fmt"

	"github.com/welldn/cribot/pkg/common"
	"github.com/welldn/cribot/pkg/database"
)

func createTestUser(db *sql.DB) error {
	testUser := common.DBUser{
		Name:     "testuser",
		Password: "testpassword",
	}

	err := database.CreateDBUser(db, testUser)
	if err != nil {
		return fmt.Errorf("failed to create test user: %v", err)
	}

	fmt.Println("Test user created successfully.")
	return nil
}
