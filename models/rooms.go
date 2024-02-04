package models

import (
	"anon/logger"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Room struct {
	Title    string `json:"title" xml:"title" form:"title" query:"title"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

func CreateRoomTable(db *sql.DB) {
	createRoomTableSQL := `
		CREATE TABLE IF NOT EXISTS room (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"title" TEXT,
			"password" TEXT,
			"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			"updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	` // SQL Statement for Create Table
	statement, err := db.Prepare(createRoomTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
		logger.LogToFile(fmt.Sprintf("Failed to create room table: %v", err))
	}
	statement.Exec() // Execute SQL Statements
	logger.LogToFile("Created Room table")
}

func InsertRoom(db *sql.DB, room *Room) (int64, error) {
	insertRoomSQL := `
		INSERT INTO room (title, password, createdAt, updatedAt)
		VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
	`

	// Execute the SQL statement
	result, err := db.Exec(insertRoomSQL, room.Title, room.Password)
	if err != nil {
		log.Fatal(err.Error())
		logger.LogToFile(fmt.Sprintf("Failed to insert room: %v", err))
		return 0, err
	}

	// Retrieve the last inserted ID
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
		logger.LogToFile(fmt.Sprintf("Failed to retrieve last inserted ID: %v", err))
		return 0, err
	}

	logger.LogToFile(fmt.Sprintf("Inserted Room with ID: %d", lastInsertedID))
	return lastInsertedID, nil
}

//WARN: this isnt part of the code, just for referencing ------------------------------------------------------------------------------------------------
type Todo struct {
	Title       string
	Description string
}

func CreateTodoTable(db *sql.DB) {
	createTodoTableSQL := `CREATE TABLE todo (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN DEFAULT FALSE,
		"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	  );` // SQL Statement for Create Table
	statement, err := db.Prepare(createTodoTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
		logger.LogToFile(fmt.Sprintf("Failed to create todo table: %v", err))
	}
	statement.Exec() // Execute SQL Statements
}

func InsertTodo(db *sql.DB, todo Todo) {
	log.Println("Inserting todo record ...")
	insertTodoSQL := `INSERT INTO todo(title, description) VALUES (?, ?)`
	statement, err := db.Prepare(insertTodoSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(todo.Title, todo.Description)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Your new todo has been created!!")
}

type TodoComplete struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// DisplayTodos retrieves and prints all todos from the database
func DisplayTodos(db *sql.DB) []TodoComplete {
	rows, err := db.Query("SELECT * FROM todo ORDER BY createdAt")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []TodoComplete
	for rows.Next() {
		var t TodoComplete
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		if err == nil {
			todos = append(todos, t)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return todos
}
