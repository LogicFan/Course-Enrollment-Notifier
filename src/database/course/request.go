package course

import (
	"net/http"
	"net/url"
)

// FetchCourseSchedule Fetch the CourseSchedule from http://www.adm.uwaterloo.ca
func FetchCourseSchedule(level string, term string, subject string, catalog string) ([]Course, error) {
	resp, err := http.PostForm(
		"http://www.adm.uwaterloo.ca/cgi-bin/cgiwrap/infocour/salook.pl",
		url.Values{
			"level":   {level},
			"sess":    {term},
			"subject": {subject},
			"cournum": {catalog},
		})
	if err != nil {
		return nil, err
	}

	return parseResponse(resp)
}

// FetchSubjectSchedule Fetch the CourseSchedule from http://www.adm.uwaterloo.ca
func FetchSubjectSchedule(level string, term string, subject string) ([]Course, error) {
	resp, err := http.PostForm(
		"http://www.adm.uwaterloo.ca/cgi-bin/cgiwrap/infocour/salook.pl",
		url.Values{
			"level":   {level},
			"sess":    {term},
			"subject": {subject},
			"cournum": {},
		})
	if err != nil {
		return nil, err
	}

	return parseResponse(resp)
}
