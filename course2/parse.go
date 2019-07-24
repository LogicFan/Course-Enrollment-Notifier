package course

import (
	"errors"
	"html"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseResponse takes a http.Resonse, return a list of course
func ParseResponse(resp *http.Response) ([]Course, error) {
	if resp == nil {
		return nil, errors.New("resp cannot be nil")
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	content := document.Selection.Children().Children().Next().Children().Next().Next().Next()
	contentText, err := content.Html()
	if err != nil {
		return nil, err
	} else if strings.Contains(contentText, "your query had no matches") {
		return nil, errors.New("Response has no matches")
	}

	courses := content.Children().Children().Children()

	retVal := make([]Course, 0, 1)
	for courses.Length() != 0 {
		// parse a Course
		var text string
		text, err = courses.Children().Html()
		if err != nil {
			courses = courses.Next()
			continue
		} else if !strings.Contains(text, "Subject") {
			courses = courses.Next()
			continue
		}
		courses = courses.Next()

		// Get subject, catalog, units and title
		subject, catalog, units, title, err := parseCourseAttr(courses)
		if err != nil {
			courses = courses.Next()
			continue
		}
		courses = courses.Next()

		// Try to get notes
		text, err = courses.Children().Html()
		notes := ""
		if err != nil {
			courses = courses.Next()
			continue
		} else if strings.Contains(text, "Notes") {
			notes, err = parseCourseNotes(courses)
			if err != nil {
				courses = courses.Next()
				continue
			}
			courses = courses.Next()
		}

		// Get sections
		parseCourseSection(courses)

		courses = courses.Next().Next()

		retVal = append(retVal,
			Course{
				subject: subject,
				catalog: catalog,
				units:   units,
				title:   title,
				notes:   notes,
			})
	}

	return retVal, nil
}

func parseCourseAttr(root *goquery.Selection) (string, string, string, string, error) {
	if root == nil {
		return "", "", "", "", errors.New("root cannot be nil")
	}

	var node *goquery.Selection
	var text string
	var err error

	node = root.Children()
	text, err = node.Html()
	if err != nil {
		return "", "", "", "", err
	}
	subject := html.UnescapeString(strings.TrimSpace(text))

	node = node.Next()
	text, err = node.Html()
	if err != nil {
		return "", "", "", "", err
	}
	catalog := html.UnescapeString(strings.TrimSpace(text))

	node = node.Next()
	text, err = node.Html()
	if err != nil {
		return "", "", "", "", err
	}
	units := html.UnescapeString(strings.TrimSpace(text))

	node = node.Next()
	text, err = node.Html()
	if err != nil {
		return "", "", "", "", err
	}
	title := html.UnescapeString(strings.TrimSpace(text))

	return subject, catalog, units, title, nil
}

func parseCourseNotes(root *goquery.Selection) (string, error) {
	if root == nil {
		return "", errors.New("root cannot be nil")
	}
	text, err := root.Children().Html()
	text = html.UnescapeString(text)
	text = strings.ReplaceAll(text, "<b>", "")
	text = strings.ReplaceAll(text, "</b>", "")
	return text, err
}

func parseCourseSection(root *goquery.Selection) (map[string]Section, error) {
	if root == nil {
		return nil, errors.New("root cannot be nil")
	}

	sections := root.Children().Next().Children().Children().Children()

	for sections.Length() != 0 {
		cla, sec, cap, erl, ins, err := parseSectionAttr(sections)
		if err != nil {
			text, _ := sections.Html()
			println(text)
		} else {
			print(cla)
			print(". ")
			print(sec)
			print(". ")
			print(cap)
			print(". ")
			print(erl)
			print(". ")
			print(ins)
			print(". ")
			print(err)
			println()
		}
		sections = sections.Next()
	}

	return nil, nil
}

func parseSectionAttr(root *goquery.Selection) (int, string, int, int, string, error) {
	if root == nil {
		return 0, "", 0, 0, "", errors.New("root cannot be nil")
	}

	var node *goquery.Selection
	var text string
	var err error

	node = root.Children()
	text, err = node.Html()
	if err != nil {
		return 0, "", 0, 0, "", err
	}
	class, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, "", 0, 0, "", err
	}

	node = node.Next()
	text, err = node.Html()
	if err != nil {
		return 0, "", 0, 0, "", err
	}
	section := strings.TrimSpace(text)

	node = node.Next().Next().Next().Next().Next()
	text, err = node.Html()
	if err != nil {
		return 0, "", 0, 0, "", err
	}
	capacity, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, "", 0, 0, "", err
	}

	node = node.Next()
	text, err = node.Html()
	if err != nil {
		return 0, "", 0, 0, "", err
	}
	enrollment, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, "", 0, 0, "", err
	}

	println(node.Length())
	println(node.Next().Length())
	println(node.Next().Next().Length())
	println(node.Next().Next().Next().Length())
	println(node.Next().Next().Next().Next().Length())
	println(node.Next().Next().Next().Next().Next().Length())
	println(node.Next().Next().Next().Next().Next().Next().Length())

	node = node.Next().Next().Next().Next().Next()
	text, err = node.Html()
	if err != nil {
		return 0, "", 0, 0, "", err
	}
	instructor := html.UnescapeString(strings.TrimSpace(text))

	return class, section, capacity, enrollment, instructor, nil
}
