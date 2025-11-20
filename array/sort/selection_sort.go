// SelectionSortArray sorts a slice of integers using the Selection Sort algorithm.
//
// Overview:
// Selection Sort divides the array into a sorted and an unsorted region. It repeatedly
// selects the smallest element from the unsorted region and swaps it with the first
// element of that region, expanding the sorted portion one element at a time.
//
// Real-world Use Case:
// - Useful for very small datasets.
// - Preferred in situations where the number of swaps must be minimized.
// - Common in teaching sorting logic due to its simple and predictable structure.
//
// When to Use:
// - The dataset is small or performance is not critical.
// - Swap operations are expensive but comparisons are cheap.
// - You need a stable and easy-to-understand sorting demonstration.
//
// When Not To Use:
// - Large datasets.
// - Performance-sensitive applications.
// - Systems with heavy comparison costs.
//
// Time Complexity:
// - Best:    O(n²)
// - Average: O(n²)
// - Worst:   O(n²)
//
// Space Complexity:
// - O(1) in-place.
//
// Algorithm Steps:
// 1. Iterate through the array.
// 2. For each position `i`, find the minimum element in the remaining unsorted region.
// 3. Swap the found minimum with the element at position `i`.
// 4. Repeat until the array is sorted.
//
// Notes:
// - Performs at most (n - 1) swaps, which can be beneficial when swap cost is high.
// - Comparison-heavy, making it slower in practice than Insertion Sort for many cases.
// - Not suitable for large or real-time applications.
//
// Related Algorithms:
// - Bubble Sort, Insertion Sort, Quick Sort, Merge Sort.

package array

func SelectionSortArray(array []int) []int {
	var currect, minIndex int
	for i := 0; i < len(array)-1; i++ {
		currect = i
		minIndex = i
		for j := i + 1; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		array[currect], array[minIndex] = array[minIndex], array[currect]
	}
	return array
}
