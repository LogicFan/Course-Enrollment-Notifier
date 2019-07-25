package database

import (
	"database/sql"

	// this is for sqlite3
	_ "github.com/mattn/go-sqlite3"

	"./course"
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

// InsertUser insert user into database
func InsertUser(user User) error {
	stmt, err := database.Prepare(`INSERT into USER_INFO 
			(email, level, term, subject, catalog, section)
		VALUES
			(?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		user.email,
		user.level,
		user.term,
		user.subject,
		user.catalog,
		user.section,
	)

	return nil
}

// UpdateCourses update section informaitons
func UpdateCourses() error {
	var query string
	query = `select DISTINCT level, term, subject FROM USER_INFO;`
	rows, err := database.Query(query)
	if err != nil {
		return err
	}

	for rows.Next() {
		var level, term, subject string
		rows.Scan(&level, &term, &subject)

		courses := course.Fe
	}

	return nil
}
