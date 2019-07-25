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

// UpdateSchedule update section informations
func UpdateSchedule() error {
	var query string
	query = `SELECT DISTINCT level, term, subject FROM USER_INFO;`
	rows, err := database.Query(query)
	if err != nil {
		return err
	}

	subjects := make([][3]string, 0, 1)

	for rows.Next() {
		var level, term, subject string
		rows.Scan(&level, &term, &subject)

		subjects = append(subjects, [3]string{level, term, subject})
	}
	rows.Close()

	database.Exec("BEGIN;")
	for _, subj := range subjects {
		println(subj[0])
		println(subj[1])
		println(subj[2])

		courses, err := course.FetchSubjectSchedule(subj[0], subj[1], subj[2])
		if err != nil {
			println(err.Error())
			continue
		}

		for _, courseObj := range courses {
			println(courseObj.ToString())
			insertCourse(subj[0], subj[1], courseObj)
		}
	}
	database.Exec("COMMIT;")

	return nil
}

func insertCourse(level string, term string, courseObj course.Course) {
	subject := courseObj.GetSubject()
	catalog := courseObj.GetCatalog()
	title := courseObj.GetTitle()
	sections := courseObj.GetSections()

	for _, sectionObj := range sections {
		class := sectionObj.GetClass()
		section := sectionObj.GetSection()
		instructor := sectionObj.GetInstructor()
		capacity := sectionObj.GetCapacity()
		enrollment := sectionObj.GetEnrollment()

		stmt, err := database.Prepare(`INSERT INTO SECTION_INFO 
			(class, level, term, subject, catalog, title, section, instructor, capacity, enrollment)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)
		defer stmt.Close()
		if err != nil {
			println(err.Error())
			continue
		}

		_, err = stmt.Exec(
			class,
			level,
			term,
			subject,
			catalog,
			title,
			section,
			instructor,
			capacity,
			enrollment,
		)
		if err != nil {
			println(err.Error())
			continue
		}
	}
}

// ClearSchedule clears the schedule
func ClearSchedule() {
	query := "DELETE FROM SECTION_INFO"
	database.Exec(query)
}

// Close close the connection
func Close() {
	database.Close()
}
