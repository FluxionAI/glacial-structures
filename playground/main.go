package main

import (
	"fmt"
	"glacial-structures/bit"
	"strconv"
)

// NumberToBinary converts an integer to its binary string.
func NumberToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

// BinaryToNumber converts a binary string to an integer.
func BinaryToNumber(bin string) (int, error) {
	v, err := strconv.ParseInt(bin, 2, 64)
	return int(v), err
}

// AlphabetToBinary converts a string to binary (8-bit per character).
func AlphabetToBinary(s string) string {
	result := ""
	for _, c := range s {
		result += fmt.Sprintf("%08b ", c)
	}
	return result
}

// BinaryToAlphabet converts 8-bit spaced binary back to string.
func BinaryToAlphabet(bin string) (string, error) {
	var out string
	var cur byte
	bits := []byte(bin)

	count := 0
	for _, b := range bits {
		if b == '0' || b == '1' {
			cur = cur<<1 | (b - '0')
			count++
			if count == 8 {
				out += string(rune(cur))
				cur = 0
				count = 0
			}
		}
	}
	if count != 0 {
		return "", fmt.Errorf("binary length not aligned to 8-bit")
	}
	return out, nil
}

func main() {
	// var z int

	x := 6 // 0001
	y := 2 // 0010
	// 00101
	// 00011
	// 00111

	e := bit.AdditionWithoutOperation(x, y)
	fmt.Println(e)

	e = bit.DivtionWithoutArithamethicOperation(x, y)
	fmt.Println(e)

	// AND operator
	// 0001 & 0010 = 0000
	// 	z = x & y
	// 	fmt.Printf("AND Operator: %d\n", z)

	// 	// OR operator
	// 	// 0001 | 0010 = 0011 (3)
	// 	z = x | y
	// 	fmt.Printf("OR Operator: %d\n", z)

	// 	// XOR operator
	// 	// 0001 ^ 0010 = 0011 (3)
	// 	z = 6 ^ 1
	// 	fmt.Printf("XOR: %d\n", z)

	// 	// LEFT SHIFT [Act as multiplier]
	// 	// x << y = 1 << 2 = 4
	// 	// 0001 -> 0100
	// 	z = 5 << 2
	// 	fmt.Printf("LEFT SHIFT: %d\n", z)

	// // RIGHT SHIFT [Act as div]
	// // x >> 1 = 1 >> 1 = 0
	// // 0001 >> 0001 = 0000
	// z = 6 >> 1
	// fmt.Printf("RIGHT SHIFT: %d\n", z)
}
