package bit_test

import (
	"testing"

	"glacial-structures/bit"
)

func TestAdditionWithoutArithmeticOperators(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{5, 7, 12},
		{0, 0, 0},
		{-1, 1, 0},
		{-5, -3, -8},
		{100, 200, 300},
	}

	for _, test := range tests {
		result := bit.AdditionWithoutArithmeticOperators(test.a, test.b)

		if result != test.expected {
			t.Errorf("AdditionWithoutArithmeticOperators (%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}
