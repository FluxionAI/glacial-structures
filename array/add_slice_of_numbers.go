package array

import (
	"fmt"
)

func AddSliceOfNumbers(a, b []int) []int {
	var aNumber, bNumber int

	// o(n)
	for i := 0; i < len(a); i++ {
		aNumber = aNumber*10 + a[i]
	}

	// o(m)
	for j := 0; j < len(b); j++ {
		bNumber = bNumber*10 + b[j]
	}

	fmt.Println("A:", aNumber, "B:", bNumber)

	// o(1)
	newValueForSlice := aNumber + bNumber
	fmt.Println("Sum:", newValueForSlice)

	// make it a fix mem by length //o(1)
	var digits []int

	// o(log(sum)) o(log(max))
	for newValueForSlice > 0 {
		// o(k square 2)
		digit := newValueForSlice % 10
		digits = append([]int{digit}, digits...)
		newValueForSlice /= 10
	}

	return digits
}

// 2242513234
//
//	012231354
//
// eg [12,23,13,54] [22,42,51,32,34]
func Addsliceofnumbers(num1, num2 []int) {
	num1, num2 = equlizer(num1, num2)
	//carry working in only + not in -
	var carry bool
	var sum int
	var carryNum int
	for i := len(num1) - 1; i > 0; i-- {
		sum = num1[i] + num2[i] + carryNum
		if sum > 100 {
			carry = true
			carryNum = sum / 100
		} else {
			carry = false
			carryNum = 0
		}
	}
	fmt.Println(sum, carry)

}

func equlizer(num1, num2 []int) ([]int, []int) {
	// return new num1 and num2 if its length is not matching
	val := abcdepth(len(num1), len(num2))
}

func abcdepth(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}
