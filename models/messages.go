package models

import (
	"anon/logger"
	"database/sql"
	"fmt"
	"log"
)

func CreateMessageTable(db *sql.DB) {
	createMessageTableSQL := `
		CREATE TABLE IF NOT EXISTS Messages (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"roomId" INTEGER NOT NULL,
			"content" TEXT NOT NULL,
			"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			"updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY ("roomId") REFERENCES "Rooms" ("id")
		);
	` // SQL Statement for Create Table
	statement, err := db.Prepare(createMessageTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
		logger.LogToFile(fmt.Sprintf("Failed to create message table: %v", err))
	}
	statement.Exec() // Execute SQL Statements
}
