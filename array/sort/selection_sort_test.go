package array

import (
	"slices"
	"testing"
)

func TestSelectionSortArray(t *testing.T) {
	tests := []struct {
		array, expected []int
	}{
		{[]int{3, 1, 5, 5, 6}, []int{1, 3, 5, 5, 6}}, // basic
		{[]int{}, []int{}},                           // empty array
		{[]int{10}, []int{10}},                       // single element
		{[]int{9, 8, 7, 6, 5}, []int{5, 6, 7, 8, 9}}, // reverse sorted
		{[]int{4, 4, 4, 4}, []int{4, 4, 4, 4}},       // all duplicates
		{[]int{-3, 10, -1, 7}, []int{-3, -1, 7, 10}}, // mix of negative & positive
	}

	for i, example := range tests {
		if got := SelectionSortArray(example.array); !slices.Equal(got, example.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i+1, example.expected, got)
		}
	}

}
