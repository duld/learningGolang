package mymath

import (
	"fmt"
	"testing"
)

// BET

// Benchmark
func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{123, 333, 444, 555, 666, 999, 101010, 1, 2, 88}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

func ExampleCenteredAvg() {
	xi := []int{123, 333, 444, 555, 666, 999, 101010, 1, 2, 88}
	r := CenteredAvg(xi)
	fmt.Println(r)
	// Output
	// 401.25
}

func TestCenteredAvg(t *testing.T) {
	xi := []int{123, 333, 444, 555, 666, 999, 101010, 1, 2, 88}
	r := CenteredAvg(xi)
	tstSum := 0.0
	for _, v := range xi[1 : len(xi)-1] {
		tstSum += float64(v)
	}
	tstSum = tstSum / float64(len(xi)-2)

	if r != tstSum {
		t.Errorf("Expected centered average sum to be %v got %v instead", tstSum, r)
	}
}
