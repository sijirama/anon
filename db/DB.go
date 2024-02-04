package db

import (
	"database/sql"
	"log"
	"os"

	"anon/logger"
	"anon/models"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var DatabaseClient *sql.DB

var dbPathUrl = "db/databse.db"

func DatabseInit() {

	_, err := os.Stat(dbPathUrl)
	if err == nil {
		//WTF: Database file exists, open it
		sqliteDatabase, err := sql.Open("sqlite3", dbPathUrl)
		if err != nil {
			log.Fatal(err)
		}
		DatabaseClient = sqliteDatabase       //INFO: let other packages use the client connection
		logger.LogToFile("Database is ready") //INFO: log success

		return
	}

	file, err := os.Create(dbPathUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	sqliteDatabase, error := sql.Open("sqlite3", dbPathUrl)
	if error != nil {
		log.Fatal(error)
	}

	models.CreateRoomTable(sqliteDatabase)
	models.CreateMessageTable(sqliteDatabase)

	DatabaseClient = sqliteDatabase //INFO: let other packages use the client connection

	logger.LogToFile("Database is ready") //INFO: log success

}
