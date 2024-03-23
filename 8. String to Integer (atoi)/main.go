package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {

	numTrimmed := strings.TrimSpace(s)

	if len(numTrimmed) == 0 {
		return 0
	}

	sign := true

	if numTrimmed[0] == '+' || numTrimmed[0] == '-' {
		sign = numTrimmed[0] != '-'
		numTrimmed = numTrimmed[1:]
	}

	if len(numTrimmed) == 0 {
		return 0
	}

	for i, v := range numTrimmed {
		if v < '0' || v > '9' {
			numTrimmed = numTrimmed[:i]
			break
		}
	}

	if len(numTrimmed) == 0 {
		return 0
	}

	//	also trim zeros for that stupid 1088 test case
	if numTrimmed[0] == '0' {
		for i, v := range numTrimmed {

			if v != '0' {
				numTrimmed = numTrimmed[i:]
				break
			}

			if i == len(numTrimmed)-1 {
				return 0
			}
		}
	}

	var result int64
	var shift int64 = 1
	const maxShift = math.MaxInt64 / 10

	for i := len(numTrimmed) - 1; i >= 0; i-- {

		posChar := int8(numTrimmed[i]) - 48
		if posChar < 0 || posChar > 9 {
			break
		}

		next := int64(posChar) * shift
		result += next
		shift = shift * 10

		if shift >= maxShift {
			result = math.MaxInt64
			break
		}
	}

	if !sign {
		result = 0 - result
	}

	if result > math.MaxInt32 {
		result = math.MaxInt32
	} else if result < math.MinInt32 {
		result = math.MinInt32
	}

	return int(result)
}

func main() {

	println(myAtoi("42"))
	println(myAtoi("422"))
	println(myAtoi("words and 987"))
	println(myAtoi("-422"))
	println(myAtoi("-91283472332"))
	println(myAtoi("4193 with words"))
	println(myAtoi("   -42"))
	println(myAtoi("20000000000000000000"))
	println(myAtoi("  0000000000012345678"))
	println(myAtoi("000000000000000000"))
}
