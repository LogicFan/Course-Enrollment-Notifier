package main

import (
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"

	"./database"
	"./email"
)

var mutex sync.Mutex

func loop(e email.Email) {
	for true {
		t := time.Now()
		println(t.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
		println("Sending email...")

		mutex.Lock()
		results := database.GetEmailList()
		mutex.Unlock()

		for _, result := range results {
			to := result.Email
			title := "Course " + result.Subject +
				" " + result.Catalog + " is Available Now"
			body := "Hi, \n" +
				"    The course " + result.Subject +
				" " + result.Catalog +
				" " + result.Section +
				" (class no. " + strconv.FormatInt(int64(result.Class), 10) +
				") with instructor " + result.Instructor +
				"now has available seat."

			content := email.Content{}
			content.Create(
				title,
				body)

			e.Send(to, content)
			println("To: " + to)
			println("Subject: " + title)
			println("Body: " + body)
			println("------------------------------")
		}

		mutex.Lock()
		for _, result := range results {
			database.DeleteUser(result.Pid)
		}
		mutex.Unlock()

		time.Sleep(
			time.Duration(30*time.Minute) - time.Since(t))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	r.ParseMultipartForm(32 << 20)

	if r.Method == "POST" {
		user := database.User{
			Email:   r.PostFormValue("email"),
			Level:   r.PostFormValue("level"),
			Term:    r.PostFormValue("term"),
			Subject: r.PostFormValue("subject"),
			Catalog: r.PostFormValue("catalog"),
			Section: r.PostFormValue("section"),
		}

		// Check if these are valid arguments
		emailRegex := true
		levelRegex, _ := regexp.Match("^(grad|under)$", []byte(user.Level))
		termRegex, _ := regexp.Match("^(1[0-9][0-9][159])$", []byte(user.Term))
		subjectRegex, _ := regexp.Match("^[A-Z]*$", []byte(user.Subject))
		catalogRegex, _ := regexp.Match("^[0-9][0-9][0-9][A-Z]*$", []byte(user.Catalog))
		sectionRegex, _ := regexp.Match("^[A-Z][A-Z][A-Z] [0-9][0-9][0-9]$", []byte(user.Section))

		if emailRegex &&
			levelRegex &&
			termRegex &&
			subjectRegex &&
			catalogRegex &&
			sectionRegex {

			println("Receive post request")
			println(user.Email + ", " +
				user.Level + ", " +
				user.Term + ", " +
				user.Subject + ", " +
				user.Catalog + ", " +
				user.Section)

			mutex.Lock()
			database.InsertUser(user)
			mutex.Unlock()

			w.Write([]byte("Success"))
			return
		}
	}

	w.Write([]byte("Failure"))
}
