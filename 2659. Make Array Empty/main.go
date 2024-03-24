package main

import (
	"sort"
)

func isSortedAscend(arr []int) bool {

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

func isSortedDescend(arr []int) bool {

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

func countOperationsToEmptyArray(nums []int) int64 {

	if isSortedAscend(nums) {
		return int64(len(nums))
	}

	if isSortedDescend(nums) {
		return int64(len(nums)) * 2
	}

	var priority []int = make([]int, len(nums))
	copy(priority, nums)
	sort.Ints(priority)

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
	//println(countOperationsToEmptyArray([]int{3, 4, -1}))
	//println(countOperationsToEmptyArray([]int{1, 2, 3}))
	//println(countOperationsToEmptyArray([]int{-20, -6}))
	println(countOperationsToEmptyArray([]int{1000000000, 999999999, 999999998, 999999997, 999999996, 999999995, 999999994, 999999993, 999999992, 999999991, 999999990, 999999989, 999999988, 999999987, 999999986, 999999985, 999999984, 999999983, 999999982, 999999981, 999999980, 999999979, 999999978, 999999977, 999999976, 999999975, 999999974, 999999973, 999999972, 999999971, 999999970, 999999969, 999999968, 999999967, 999999966, 999999965, 999999964, 999999963, 999999962, 999999961, 999999960, 999999959, 999999958, 999999957, 999999956, 999999955, 999999954, 999999953, 999999952, 999999951, 999999950, 999999949, 999999948, 999999947, 999999946, 999999945, 999999944, 999999943, 999999942, 999999941, 999999940, 999999939, 999999938, 999999937, 999999936, 999999935, 999999934, 999999933, 999999932, 999999931, 999999930, 999999929, 999999928, 999999927, 999999926})) //	2851 ops
}
