package course

import (
	"net/http"
	"net/url"
)

// FetchCourseSchedule Fetch the CourseSchedule from http://www.adm.uwaterloo.ca
func FetchCourseSchedule(term string, level string, subject string, catalog string) (*http.Response, error) {
	return http.PostForm(
		"http://www.adm.uwaterloo.ca/cgi-bin/cgiwrap/infocour/salook.pl",
		url.Values{
			"level":   {level},
			"sess":    {term},
			"subject": {subject},
			"cournum": {catalog},
		})
}

// FetchSubjectSchedule Fetch the CourseSchedule from http://www.adm.uwaterloo.ca
func FetchSubjectSchedule(term string, level string, subject string) (*http.Response, error) {
	return http.PostForm(
		"http://www.adm.uwaterloo.ca/cgi-bin/cgiwrap/infocour/salook.pl",
		url.Values{
			"level":   {level},
			"sess":    {term},
			"subject": {subject},
			"cournum": {},
		})
}
