package array

func BubbleSortArray(array []int) []int {
	n := len(array)
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap = true
			}
		}
		if swap {
			break
		}
	}
	return array
}
