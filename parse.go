package aocutils

import "strconv"

// Returns a number which was parsed from the given string. Uses strconv.Atoi. If the number cannot be parsed, returns an error.
func ParseNum(s string) (int, error) {
	n, err := strconv.Atoi(s)
	return n, err
}

// Returns a number which was parsed from the given string or panics if it fails. Uses strconv.Atoi.
func ParseNumOrPanic(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("cannot convert to int")
	}

	return n
}
