package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {

	isNegative := x < 0
	numberAsString := strconv.Itoa(x)

	if isNegative {
		numberAsString = numberAsString[1:]
	}

	var reversed string
	if isNegative {
		reversed += "-"
	}

	for i := len(numberAsString) - 1; i >= 0; i-- {
		reversed += string(numberAsString[i])
	}

	parsed, err := strconv.ParseInt(reversed, 10, 64)
	if err != nil {
		return 0
	}

	if parsed > math.MaxInt32 || parsed < math.MinInt32 {
		return 0
	}

	return int(parsed)
}

func main() {

	println(reverse(123))
	println(reverse(-123))
	println(reverse(120))

}
