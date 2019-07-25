package main

import (
	"net/http"
	"os"

	"./database"
	"./email"
)

func main() {
	var err error

	argv := os.Args
	if len(argv) != 2 {
		return
	}

	// init database
	err = database.Init(argv[1] + "/config/info.db")
	defer database.Close()
	if err != nil {
		println(err.Error())
		return
	}

	// init email
	e := email.Email{}
	err = e.InitByFile(argv[1] + "/config/email.config")
	if err != nil {
		println(err.Error())
		return
	}

	go loop(e)

	mux := http.NewServeMux()
	mux.HandleFunc("/notifier", handler)
	http.ListenAndServe(":666", mux)
}
