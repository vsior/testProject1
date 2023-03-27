//go:build integration
// +build integration

package adder

import "testing"

func Test_TableTests(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"addition", 2, 2, 4},
		{"subtraction", 20, 20, 40},
		{"multiplication", 22, 22, 44},
		{"division", 21, 21, 420},
		{"bad_division", 2, 0, 2},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := addNumbers(d.num1, d.num2)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
		})
	}
}

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
