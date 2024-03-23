package main

import (
	"fmt"
	"math"
	"strconv"
)

func parseBuffer(str string, width int64) []byte {

	byteWidth := width / 8
	if (width % 8) > 0 {
		byteWidth++
	}

	result := make([]byte, byteWidth)

	posShift := len(str) % 8

	for i := 0; i < len(str); i++ {

		idx := i + posShift
		byteIdx := idx / 8
		bitIdx := (idx % 8)

		var setBit int8
		if str[i] == '1' {
			setBit = 1
		}

		println("bitIdx", bitIdx, setBit, string(str[i]))

		result[byteIdx] |= byte(setBit << bitIdx)
	}

	return result
}

func printByteAsBits(data []byte) {
	for _, b := range data {
		for i := 7; i >= 0; i-- {
			bit := (b >> uint(i)) & 1
			fmt.Print(strconv.FormatInt(int64(bit), 2))
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func addBinary(a string, b string) string {

	bitWidth := math.Max(float64(len(a)), float64(len(b)))

	printByteAsBits(parseBuffer(a, int64(bitWidth)))
	printByteAsBits(parseBuffer(b, int64(bitWidth)))

	aNum, _ := strconv.ParseInt(a, 2, 32)
	bNum, _ := strconv.ParseInt(b, 2, 32)
	return strconv.FormatInt(aNum+bNum, 2)
}

func main() {
	//addBinary("10", "1")
	//addBinary("11", "1")
	//addBinary("1010", "1011")
	//println(addBinary("1010", "1011"))
}
