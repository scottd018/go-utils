package pointers

// String returns the pointer value of a string.
func String(str string) *string {
	return &str
}

// Integer returns the pointer value of an integer.
func Integer(i int) *int {
	return &i
}

// Bool returns the pointer value of a bool.
func Bool(b bool) *bool {
	return &b
}
