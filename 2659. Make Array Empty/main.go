package main

import (
	"fmt"
	"sort"
)

func createPriority(nums []int) []int {

	var priority []int = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		priority[i] = i
	}

	sort.Slice(priority, func(x, y int) bool {
		return nums[x] < nums[y]
	})

	fmt.Printf("%v\n", priority)

	return priority
}

func countOperationsToEmptyArray(nums []int) int64 {

	priority := createPriority(nums)

	temp := len(nums)
	var steps int64 = 0

	for i := 1; i < len(nums); i++ {
		if priority[i] < priority[i-1] {
			steps += int64(temp)
			temp = len(nums) - i
		}
	}

	steps += int64(temp)
	return steps
}

func main() {
	println(countOperationsToEmptyArray([]int{3, 4, -1}))
	println(countOperationsToEmptyArray([]int{1, 2, 3}))
	println(countOperationsToEmptyArray([]int{-20, -6}))
}
