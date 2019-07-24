package main

import "./email"

func main() {
	e := email.TestEmail
	content := email.Content{}
	err := content.Create(
		"fanyongda2012@gmail.com",
		"CS 135 LEC001 Has Spot Now",
		"The course CS 135 LEC 001, now has a spot, please enroll")

	err = e.Send("fanyongda2012@gmail.com", content)
	if err != nil {
		println(err.Error())
	}
}
