package course

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Section contains a capacity and enrollment
type Section struct {
	capacity   int
	enrollment int
}

// ParseSection parse from html parse tree to a Section object
// @ root: an html parse tree
func ParseSection(root *goquery.Selection) (string, Section, error) {
	sectionAttrs := make([]string, 0, 13)

	i := 0

	// convert html to an array
	for node := root.Children(); node.Length() != 0; node = node.Next() {
		text, err := node.Html()

		if err != nil {
			return "", Section{}, err
		}

		print(i)
		print(": \t")
		println(text)

		i++

		sectionAttrs = append(sectionAttrs, strings.TrimSpace(text))
	}

	// 0: Class
	// 1: Comp Sec
	// 2: Comp Loc
	// 3: Assoc. Class
	// 4: Rel 1
	// 5: Rel 2
	// 6: Enrl Cap
	// 7: Enrl Tot
	// 8: Wait Cap
	// 9: Wait Tot
	// 10: Time Days/Date
	// 11: BldgeRoom
	// 12: Instructor
	// 13: NIL

	if len(sectionAttrs) != 13 {
		return "", Section{}, errors.New("section: wrong number of attribute")
	}

	cap, err := strconv.Atoi(sectionAttrs[6])
	if err != nil {
		return "", Section{}, err
	}

	enrl, err := strconv.Atoi(sectionAttrs[7])
	if err != nil {
		return "", Section{}, err
	}

	sectionObject := Section{capacity: cap, enrollment: enrl}
	sectionName := sectionAttrs[1]

	return sectionName, sectionObject, nil
}

// Capacity return the capacity of a section
func (sec Section) Capacity() int {
	return sec.capacity
}

// Enrollment return the enrollment number of a section
func (sec Section) Enrollment() int {
	return sec.enrollment
}

// ToString return a string of object a section
func (sec Section) ToString() string {
	return "Enrl Cap: " + strconv.Itoa(sec.Capacity()) + "Enrl Tot:" + strconv.Itoa(sec.Enrollment())
}

// Course contains a map from section name to Section
type Course struct {
	sections map[string]Section
}

// ParseCourse return a course of given html parse tree
func ParseCourse(root *goquery.Selection) (Course, error) {
	node := root.Children().Next().Children().Children().Children()

	sectionMap := make(map[string]Section)

	for node = node.Next(); node.Length() != 0; node = node.Next() {
		name, sec, err := ParseSection(node)

		if err == nil {
			sectionMap[name] = sec
		}
	}

	if len(sectionMap) == 0 {
		return Course{}, errors.New("no section found, this course may not exists")
	}

	return Course{sections: sectionMap}, nil
}

// Sections asjdg;j;laskdgflkas
func (cou Course) Sections() map[string]Section {
	return cou.sections
}

// ToString return a string of object a section
func (cou Course) ToString() string {
	retVal := ""
	for name, obj := range cou.sections {
		retVal = retVal + name + ": " + obj.ToString() + "\n"
	}
	return retVal
}
