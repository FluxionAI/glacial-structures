package bit

import "fmt"

// 3 ,5

// / 0011
// / 0101
// / 0001 carry

// / 0011
// / 0101
// / 0111 carry

// b = 5
//

func AdditionWithoutOperation(a, b int) int {
	// for b != 0 {
	// 	carry := a & b
	// 	a = a ^ b
	// 	// Obtain carry and shift it left
	// 	b = carry << 1
	// }

	sum := a & b
	// 00111 ^
	// 00001 &
	fmt.Println(sum)
	return a
}
