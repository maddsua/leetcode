package main

import (
	"fmt"
)

func addBinary(a string, b string) string {

	bitWidth := len(a)
	if len(b) > bitWidth {
		bitWidth = len(b)
		a = fmt.Sprintf("%*s", bitWidth, a)
	} else if len(b) < bitWidth {
		b = fmt.Sprintf("%0*s", bitWidth, b)
	}

	var carryflag bool
	var temp []byte

	for i := bitWidth - 1; i >= 0; i-- {

		var sum int
		if carryflag {
			sum = 1
		}

		if a[i] == '1' {
			sum++
		}
		if b[i] == '1' {
			sum++
		}

		if sum < 2 {

			if sum == 0 {
				temp = append(temp, '0')
			} else {
				temp = append(temp, '1')
			}

			carryflag = false

		} else {
			temp = append(temp, '0')
			carryflag = true
		}
	}

	if carryflag {
		temp = append(temp, '1')
	}

	var result []byte

	for i := len(temp) - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}

	return string(result)
}

func main() {
	//println(addBinary("10", "1"))
	//println(addBinary("11", "1"))
	println(addBinary("1010", "1011"))
	//println(addBinary("1010", "1011"))
}
