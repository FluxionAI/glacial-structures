package array

import (
	"slices"
	"testing"
)

func TestBubbleSortArray(t *testing.T) {
	tests := []struct {
		array, expected []int
	}{
		{[]int{3, 1, 5, 5, 6}, []int{1, 3, 5, 5, 6}},
	}

	for i, example := range tests {
		if got := BubbleSortArray(example.array); !slices.Equal(got, example.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, example.expected, got)
		}
	}
}
