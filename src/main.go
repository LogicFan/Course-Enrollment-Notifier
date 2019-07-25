package main

import (
	"net/http"

	"./database"
	"./email"
)

func main() {
	var err error

	// init database
	err = database.Init("../config/info.db")
	defer database.Close()
	if err != nil {
		println(err.Error())
		return
	}

	// init email
	e := email.Email{}
	err = e.InitByFile("../config/email.config")
	if err != nil {
		println(err.Error())
		return
	}

	go loop(e)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
