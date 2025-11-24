package bit

func DivtionWithoutArithamethicOperation(x, y int) int {
	quotient := 0
	power := uint64(32)
	yPower := y << power
	reminder := x

	for reminder >= y {
		for yPower > reminder {
			yPower >>= 1
			power--
		}
		quotient += 1 << power
		reminder -= yPower
	}
	return quotient
}
