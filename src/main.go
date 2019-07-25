package main

import "./database"

func main() {
	err := database.Init("../config/info.db")
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

	database.ClearSchedule()

	database.UpdateSchedule()
}
