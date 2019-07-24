package course

// Section contains information about sections
type Section struct {
	class      int
	section    string
	capacity   int
	enrollment int
	instructor string
	reserves   []Reserve
	held       string
}

// Reserve contains information about reserve
type Reserve struct {
	condition  string
	capacity   int
	enrollment int
}
