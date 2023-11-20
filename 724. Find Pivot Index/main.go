package main

import "fmt"

func pivotIndex(nums []int) int {

	left := 0
	right := 0

	for _, v := range nums {
		right += v
	}

	for i, v := range nums {

		left += v
		right -= v

		if left >= right {
			return i
		}
	}

	return 0
}


func main() {

	var data = []int{1, 7, 3, 6, 5, 6}

	fmt.Printf("%d\n", pivotIndex(data))

}
