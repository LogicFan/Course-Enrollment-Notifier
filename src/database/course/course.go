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

// GetSubject return subject
func (course Course) GetSubject() string {
	return course.subject
}

// GetCatalog return subject
func (course Course) GetCatalog() string {
	return course.catalog
}

// GetTitle return subject
func (course Course) GetTitle() string {
	return course.title
}

// GetSections return sections
func (course Course) GetSections() map[string]Section {
	return course.sections
}
