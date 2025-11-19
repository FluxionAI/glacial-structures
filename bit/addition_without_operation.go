package bit

import "fmt"

// 3 ,5
func AdditionWithoutOperation(a, b int) int {
	for b != 0 {
		carry := a & b
		fmt.Println("carry:", carry, " a:", a, " b:", b)
		a = a ^ b
		// Obtain carry and shift it left
		b = carry << 1
		fmt.Println("b after carry shift:", b, "carry:", carry, " a:", a)
	}
	return a
}
