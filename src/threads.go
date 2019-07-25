package main

import (
	"net/http"
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
				" " + result.Catalog + " Has Spot Now"
			body := "Hi, \n" +
				"    The course " + result.Subject +
				" " + result.Catalog +
				" " + result.Section +
				" (class no. " + strconv.FormatInt(int64(result.Class), 10) +
				") with instructor " + result.Instructor +
				"now has spot."

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
			time.Duration(5*time.Minute) - time.Since(t))
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
	}
}
