/*
Create a dog package. The dog package should have an exported func
“Years” which takes human years and turns them into dog years
(1 human year = 7 dog years). Document  your code with comments.
Use this code in func main.

run your program and make sure it works
run a local server with godoc and look at your documentation.
*/

// Package dog is a collection of functions and data structures
// associated with dogs and dog releated tasks.
package dog

// Years converts the number of years from human years, to dog years.
func Years(years int) int {
	return years * 7
}
