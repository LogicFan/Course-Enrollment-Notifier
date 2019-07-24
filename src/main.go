package main

import "./email"

func main() {
	e := email.TestEmail
	err := e.Send("notifier.uwaterloo@gmail.com", []string{"fanyongda2012@hotmail.com"}, []byte("Subject: Hello world \n Hello world Again"))
	if err != nil {
		println(err.Error())
	}
}
