package word

import (
	"testing"

	"github.com/duld/learningGolang/exercises/learning-golang/level_13/ex_1302/quote"
)

type testCount struct {
	value  string
	length int
}

func TestUseCount(t *testing.T) {
	s := "one two two three three three four four four four"
	uc := UseCount(s)

	ab := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	for k, v := range uc {
		if uc[k] != ab[k] {
			t.Errorf("Expected word: %v to be counted %v number of times", k, v)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []testCount{
		{"one two three four, five. six seven eight nine ten.", 10},
		{"                                    ", 0},
		{"hello, it's nice to meet you,", 6},
		{"      this is an annoying      string      ", 5},
	}

	var res int
	for _, s := range tests {
		res = Count(s.value)
		if res != s.length {
			t.Errorf("Expected string: \n%v\nTo have %v number of words got %v instead", s.value, s.length, res)
		}
	}
}

func BenchmarkCount10(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkCount100(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Count(quote.SunAlso)
	}
}
