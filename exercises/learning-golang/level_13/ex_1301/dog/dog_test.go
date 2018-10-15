package dog

import (
	"fmt"
	"testing"
)

// BET - BENCHMARK; EXAMPLE; TEST
// BENCHMARKS
// BENCHMARK - YEARS 10; 1,000; 100,000;
func BenchmarkYears10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYears1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(1000)
	}
}

func BenchmarkYears100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(100000)
	}
}

// BENCHMARK - YEARS 10; 1,000; 100,000;
func BenchmarkYearsTwo10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func BenchmarkYearsTwo1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(1000)
	}
}

func BenchmarkYearsTwo100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(100000)
	}
}

// EXAMPLES
func ExampleYears() {
	fmt.Println(Years(10))
	// Output: 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output: 70
}

// TESTS
func TestYears(t *testing.T) {

	for i := 10; i < 100; i++ {
		dogAge := Years(i)
		if dogAge != i*7 {
			t.Errorf("Passed in %v expected %v got %v instead", i, i*7, dogAge)
		}
	}
}

func TestYearsTwo(t *testing.T) {

	for i := 10; i < 100; i++ {
		dogAge := YearsTwo(i)
		if dogAge != i*7 {
			t.Errorf("Passed in %v expected %v got %v instead", i, i*7, dogAge)
		}
	}
}
