package main

import (
	"./database"
	"./email"
)

func main() {
	err := database.Init("../config/info.db")
	defer database.Close()
	if err != nil {
		println(err.Error())
	}

	database.InsertUser(database.User{
		Email:   "fanyongda2012@hotmail.com",
		Level:   "under",
		Term:    "1199",
		Subject: "CS",
		Catalog: "135",
		Section: "LEC 001",
	})

	database.InsertUser(database.User{
		Email:   "fanyongda2012@hotmail.com",
		Level:   "under",
		Term:    "1199",
		Subject: "MATH",
		Catalog: "135",
		Section: "LEC 001",
	})

	e := email.Email{}
	err = e.InitByFile("../config/email.config")
	if err != nil {
		println(err.Error())
		return
	}

	loop(e)
}
