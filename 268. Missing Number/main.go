package main

import (
	"sort"
)

func missingNumber(nums []int) int {

	sort.Ints(nums)

	lastNum := 0
	for _, v := range nums {
		if v-lastNum > 1 {
			return lastNum + 1
		}
		lastNum = v
	}

	if nums[0] == 0 {
		return lastNum + 1
	}

	return 0
}


func main() {

	res := missingNumber([]int{1, 2, 3, 4, 8})
	print(res)
}
