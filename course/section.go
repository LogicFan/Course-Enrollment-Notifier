package course

import "strconv"

// Section contains information about sections
type Section struct {
	class      int
	section    string
	capacity   int
	enrollment int
	instructor string
}

// ToString returns the string of given object
func (section Section) ToString() string {
	return strconv.FormatInt(int64(section.class), 10) +
		", " + section.section +
		", cap: " + strconv.FormatInt(int64(section.capacity), 10) +
		", erl: " + strconv.FormatInt(int64(section.enrollment), 10) +
		", inst: " + section.instructor
}
