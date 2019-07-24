package course

import (
	"errors"
	"net/http"
	"strings"
)

// ParseResponse takes a http.Resonse, return a list of course
func ParseResponse(resp *http.Response) ([]Course, error) {
	if resp == nil {
		return nil, errors.New("resp cannot be nil")
	}

	document, err := ParseHTML(resp.Body)
	if err != nil {
		return nil, err
	}

	var text string
	contentNode := document.Children().Children().Next().Children().Next().Next().Next()

	text = contentNode.Children().Next().Children().Lexeme()
	if strings.Contains(text, "Sorry, but your query had no matches.") {
		return nil, errors.New("Response has no matches")
	}

	coursesNode := contentNode.Children().Children().Children()
	retVal := make([]Course, 0, 1)
	for !coursesNode.Nil() {
		course := Course{}

		coursesNode, err = course.parseInit(coursesNode)
		if err != nil {
			coursesNode = coursesNode.Next()
			continue
		}

		coursesNode, err = course.parseAttr(coursesNode)
		if err != nil {
			coursesNode = coursesNode.Next()
			continue
		}

		coursesNode, err = course.parseNote(coursesNode)
		if err != nil {
			coursesNode = coursesNode.Next()
			continue
		}

		coursesNode, err = course.parseSection(coursesNode)
		if err != nil {
			coursesNode = coursesNode.Next()
			continue
		}

		retVal = append(retVal, course)
	}

	return retVal, nil
}

func (course *Course) parseInit(root Node) (Node, error) {
	node := root

	for !node.Nil() {
		text := node.Children().Children().Text()
		if strings.Contains(text, "Subject") {
			return node.Next(), nil
		}
		node = node.Next()
	}

	return root, errors.New("no course html find")
}

func (course *Course) parseAttr(root Node) (Node, error) {
	if course == nil {
		return root, errors.New("course cannot be nil")
	}

	node := root.Children()
	course.subject = node.Text()

	node = node.Next()
	course.catalog = node.Text()

	node = node.Next()
	course.units = node.Text()

	node = node.Next()
	course.title = node.Text()

	return root.Next(), nil
}

func (course *Course) parseNote(root Node) (Node, error) {
	if course == nil {
		return root, errors.New("course cannot be nil")
	}

	node := root.Children().Children()
	text := node.Text()
	if strings.Contains(text, "Notes") {
		course.notes = node.Next().Text()
		return root.Next(), nil
	}

	return root, nil
}

func (course *Course) parseSection(root Node) (Node, error) {
	if course == nil {
		return root, errors.New("course cannot be nil")
	}

	node := root.Children().Next().Children().Children().Children()
	text := node.Text()
	println(text)

	return root, nil
}

func (section *Section) parseInit(root Node) (Node, error) {
	return root, nil
}