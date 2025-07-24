package main

import (
	"fmt"
	"testing"
)

// Simple functions to test
func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Test functions
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"zero and positive", 0, 5, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"zero multiplication", 5, 0, 0},
		{"negative numbers", -2, -3, 6},
		{"mixed signs", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"even positive", 4, true},
		{"odd positive", 5, false},
		{"zero", 0, true},
		{"even negative", -4, true},
		{"odd negative", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.n)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "Hello, 世界", "界世 ,olleH"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(123, 456)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(123, 456)
	}
}

func BenchmarkReverse(b *testing.B) {
	input := "Hello, World! This is a benchmark test."
	for i := 0; i < b.N; i++ {
		Reverse(input)
	}
}

// Example test (shows in documentation)
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

func ExampleReverse() {
	result := Reverse("hello")
	fmt.Println(result)
	// Output: olleh
}

/*
Running the tests:

# Run all tests
go test

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestAdd

# Run benchmarks
go test -bench=.

# Run benchmarks with memory stats
go test -bench=. -benchmem

# Check test coverage
go test -cover

# Generate detailed coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

Example output:
=== RUN   TestAdd
=== RUN   TestAdd/positive_numbers
=== RUN   TestAdd/zero_and_positive
=== RUN   TestAdd/negative_numbers
=== RUN   TestAdd/mixed_signs
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/positive_numbers (0.00s)
    --- PASS: TestAdd/zero_and_positive (0.00s)
    --- PASS: TestAdd/negative_numbers (0.00s)
    --- PASS: TestAdd/mixed_signs (0.00s)
PASS
ok      go-learning-guide       0.002s
*/
