package main

import "./database"

func main() {
	err := database.Init("../config/info.db")
	if err != nil {
		println(err.Error())
	}
}
