package database

// Result result of email list
type Result struct {
	Pid        int
	Email      string
	Subject    string
	Catalog    string
	Section    string
	Title      string
	Instructor string
	Class      int
}

// GetEmailList return a list of result, need lock
func GetEmailList() []Result {
	clearSchedule()
	updateSchedule()

	query := `SELECT u.pid,
		u.email, 
		s.subject, 
		s.catalog, 
		s.section, 
		s.title, 
		s.instructor,
		s.class
	FROM USER_INFO u INNER JOIN SECTION_INFO s
	ON u.level = s.level 
		AND u.term = s.term 
		AND u.subject = s.subject 
		AND u.catalog = s.catalog 
		AND u.section = s.section
	WHERE s.enrollment < s.capacity;`

	retVal := make([]Result, 0, 1)

	rows, err := database.Query(query)
	defer rows.Close()
	if err != nil {
		println(err.Error())
		return retVal
	}

	for rows.Next() {
		result := Result{}
		rows.Scan(&result.Pid,
			&result.Email,
			&result.Subject,
			&result.Catalog,
			&result.Section,
			&result.Title,
			&result.Instructor,
			&result.Class)

		retVal = append(retVal, result)
	}

	return retVal
}

// DeleteUser delete user by pid
func DeleteUser(pid int) {
	stmt, err := database.Prepare(`DELETE FROM USER_INFO WHERE pid = ?;`)
	defer stmt.Close()
	if err != nil {
		println(err.Error())
	}

	_, err = stmt.Exec(pid)
	if err != nil {
		println(err.Error())
	}
}
