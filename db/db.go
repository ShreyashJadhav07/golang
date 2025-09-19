package db

import (
	"database/sql"

   _ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB
// this function is used for initialize or make connextion to data base
func InitDB(){
	var err error
	
	DB,err =sql.Open("sqlite","api.db")
	// to open the connection with data base
	if err!= nil{
		panic("Could not connct to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	//   how much conextion we need to open if no one using thic connection

	createTables()
}

func createTables(){
	createEventTables:= `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER

	)
	`
	_, err := DB.Exec(createEventTables)

	if err!=nil{
		panic("Could not create events table.")
	}
}

