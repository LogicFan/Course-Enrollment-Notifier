package main

import (
	"fmt"

	"./course"
)

func main() {
	fmt.Println("Sending Post Request!")
	resp, _ := course.FetchSubjectSchedule("1199", course.Undergraduate, "CS")

	ret, _ := course.ParseResponse(resp)

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
