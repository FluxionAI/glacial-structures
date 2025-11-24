package array

// if i have slice of [99] [9,9].
// what will be the sum ? in which decimal i need to set.
func AddSliceOfTwoNumbers(num1, num2 []int) []int {
	num1, num2 = equlizer(num1, num2)

	carry := false

	for i := len(num1) - 1; i > -1; i-- {
		num1[i] += num2[i]
		if carry {
			num1[i]++
			carry = false
		}
		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}
	if carry {
		num1 = append([]int{1}, num1...)
	}

	return num1
}

func equlizer(num1, num2 []int) ([]int, []int) {
	balanceLenght := abcdiff(len(num1), len(num2))
	if balanceLenght != 0 {
		razer := make([]int, balanceLenght)
		if len(num1) > len(num2) {
			num2 = append(razer, num2...)
		} else {
			num1 = append(razer, num1...)
		}
	}
	return num1, num2
}

func abcdiff(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}
