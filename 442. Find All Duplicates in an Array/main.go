package main

import "fmt"

func findDuplicates(nums []int) []int {

	vset := map[int]bool{}
	result := []int{}

	for _, val := range nums {

		_, has := vset[val]
		if has {
			result = append(result, val)
			continue
		}

		vset[val] = true
	}

	return result
}

func main() {

	fmt.Printf("%v\n", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Printf("%v\n", findDuplicates([]int{1, 1, 2}))

}
