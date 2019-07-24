package course

// level of the course
const (
	Undergraduate = "under"
	Graduate      = "grad"
)

// Course a class of a course, contains all section and Notes
type Course struct {
	subject  string
	catalog  string
	units    string
	title    string
	notes    string
	sections map[string]Section
}

// ToString return a string of the object
func (course Course) ToString() string {
	retVal := course.subject +
		" " + course.catalog +
		" " + course.title +
		" \tUnits: " + course.units +
		" \tNotes: " + course.notes
	for k, v := range course.sections {
		retVal = retVal + "\n\t[" + k + "] " + v.ToString()
	}
	return retVal
}