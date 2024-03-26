package main

import (
	"sort"
)

func firstMissingPositive(nums []int) int {

	sort.Ints(nums)

	var smallest int = 0
	var nextSmallest int = 0
	var biggest int = 0

	for _, v := range nums {

		if v > smallest {

			nextSmallest = smallest
			smallest = v

			if smallest-nextSmallest > 1 {
				break
			}

		}

		if v > biggest {
			biggest = v
		}
	}

	if biggest == smallest {
		return biggest + 1
	}

	return nextSmallest + 1
}

func main() {
	println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
