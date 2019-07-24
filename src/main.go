package main

import "./email"

func main() {
	e := email.Email{}
	err := e.InitByFile("../config/email.config")
	if err != nil {
		print(err.Error())
		return
	}
	content := email.Content{}
	content.Create(
		"CS 135 LEC001 Has Spot Now",
		"The course CS 135 LEC 001, now has a spot, please enroll")

	err = e.Send("fanyongda2012@hotmail.com", content)
	if err != nil {
		println(err.Error())
	}
}
