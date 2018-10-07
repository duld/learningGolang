/*
We are going to learn about testing next. For this exercise, take a moment and
see how much you can figure out about testing by reading the testing
documentation & also Caleb Doxsey’s article on testing. See if you can get a
basic example of testing working.
*/
package stringy

import (
	"testing"
)

func TestStringy(t *testing.T) {
	a := 33
	b := ToString(a)
	_, ok := interface{}(b).(string)
	if !ok {
		t.Error("Expected int to be coverted to a string")
	}
}
