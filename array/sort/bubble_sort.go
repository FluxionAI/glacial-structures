// BubbleSortArray sorts a slice of integers using the Bubble Sort algorithm.
//
// Overview:
// Bubble Sort works by repeatedly comparing adjacent elements and swapping them
// if they are in the wrong order. This implementation includes an early-stop
// optimization using a `swapped` flag, which exits as soon as the array becomes
// fully sorted.
//
// Real-world Use Case:
// - Suitable only for very small or nearly sorted datasets.
// - Commonly used for teaching sorting fundamentals.
// - Rarely appropriate for production systems due to poor performance.
//
// When to Use:
// - Input size is small.
// - Readability and simplicity matter more than performance.
// - Data is already nearly sorted.
//
// When Not to Use:
// - Input size is moderate or large.
// - Performance is a priority.
// - Systems requiring efficient sorting under load.
//
// Time Complexity:
// - Best:    O(n)      // Achieved due to early break when no swaps occur.
// - Average: O(n^2)
// - Worst:   O(n^2)
//
// Space Complexity:
// - O(1) in-place.
//
// Algorithm Steps:
// 1. Loop through the array multiple times.
// 2. Compare adjacent pairs and swap if out of order.
// 3. Stop early if a full pass results in no swaps.
//
// Notes:
// - `swapped` avoids unnecessary passes when the array is already sorted.
// - Inefficient for large datasets compared to QuickSort, MergeSort, or HeapSort.
//
// Related Algorithms:
// - Selection Sort, Insertion Sort, Merge Sort, Quick Sort.

package array

func BubbleSortArray(array []int) []int {
	n := len(array)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return array
}
