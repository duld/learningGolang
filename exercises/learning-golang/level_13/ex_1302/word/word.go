package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
// UseCount counts how many times each word occurs in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the number of words in a string
func Count(s string) int {
	// handle empty string
	c := strings.Fields(s)

	return len(c)
}
