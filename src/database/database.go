package database

import (
	"database/sql"

	// this is for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

// Init initialize the sqlite
func Init(path string) error {
	var err error
	database, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	var query string
	// create table for userinfo
	query = `CREATE TABLE IF NOT EXISTS USER_INFO(
		pid 	INTEGER 		PRIMARY KEY AUTOINCREMENT,
		email 	VARCHAR(100) 	NOT null,
		level 	VARCHAR(20) 	NOT null,
		term 	CHAR(4) 		NOT null,
		subject VARCHAR(10) 	NOT null,
		catalog VARCHAR(10) 	NOT null,
		section VARCHAR(10) 	NOT null,
		UNIQUE(email, level, term, subject, catalog, section)
	);`

	_, err = database.Exec(query)
	if err != nil {
		return err
	}

	// create table for sections
	query = `CREATE TABLE IF NOT EXISTS SECTION_INFO(
		class 		INTEGER 	PRIMARY KEY,
		level 		VARCHAR(20) NOT null,
		term 		CHAR(4) 	NOT null,
		subject 	VARCHAR(10) NOT null,
		catalog 	VARCHAR(10) NOT null,
  		title 		VARCHAR(50) NOT null,
		section 	VARCHAR(10) NOT null,
  		instructor 	VARCHAR(50) NOT null,
  		capacity 	INTEGER 	NOT null,
  		enrollment 	INTEGER 	NOT null
	);`

	_, err = database.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// InsertUser insert user into database, need lock
func InsertUser(user User) error {
	stmt, err := database.Prepare(`INSERT INTO USER_INFO 
			(email, level, term, subject, catalog, section)
		VALUES
			(?, ?, ?, ?, ?, ?);`)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		user.Email,
		user.Level,
		user.Term,
		user.Subject,
		user.Catalog,
		user.Section,
	)

	return nil
}

// Close close the connection
func Close() {
	database.Close()
}
