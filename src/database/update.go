package database

import "./course"

func updateSchedule() error {
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

		courses, err := course.FetchSubjectSchedule(subj[0], subj[1], subj[2])
		if err != nil {
			println(err.Error())
			continue
		}

		for _, courseObj := range courses {
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

func clearSchedule() {
	query := "DELETE FROM SECTION_INFO"
	database.Exec(query)
}
