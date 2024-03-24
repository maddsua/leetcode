package main

import (
	"sort"
)

func prioritySort(nums []int) []int {

	var priority []int = make([]int, len(nums))
	copy(priority, nums)

	sort.Slice(priority, func(i, j int) bool {
		return priority[i] < priority[j]
	})

	return priority
}

func countOperationsToEmptyArray(nums []int) int64 {

	priority := prioritySort(nums)

	var steps int64
	for ; len(nums) > 1; steps++ {

		if priority[0] == nums[0] {
			nums = nums[1:]
			priority = priority[1:]
		} else {
			temp := nums[0]
			nums = nums[1:]
			nums = append(nums, temp)
		}
	}

	return steps + 1
}

func main() {
	println(countOperationsToEmptyArray([]int{3, 4, -1}))
	println(countOperationsToEmptyArray([]int{1, 2, 3}))
	println(countOperationsToEmptyArray([]int{-20, -6}))
}
