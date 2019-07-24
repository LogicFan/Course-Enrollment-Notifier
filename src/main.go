package main

import (
	"fmt"

	"./course"
)

func main() {
	fmt.Println("Sending Post Request!")
	ret, _ := course.FetchSubjectSchedule("1199", course.Undergraduate, "CS")

	for _, x := range ret {
		println(x.ToString())
	}
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// courses, err := course.ParseResponse(resp)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// for _, c := range courses {
	// 	println(c.ToString())
	// }
}
