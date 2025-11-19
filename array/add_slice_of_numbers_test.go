package array

import (
	"testing"
)

func TestAddSliceOfNumbers(t *testing.T) {
	tests := []struct {
		num1, num2, sum []int
	}{
		{[]int{2}, []int{4}, []int{6}},
	}

	for i, test := range tests {
		expected := AddSliceOfNumbers(test.num1, test.num2)
		if len(expected) != len(test.sum) {
			t.Errorf("addSliceOperation of %d is not matching", i)
		}
		for j := 0; j < len(expected); j++ {
			if expected[i] != test.sum[i] {
				t.Errorf("addSliceOperation of %d is not matching", i)
			}
		}
	}
}
