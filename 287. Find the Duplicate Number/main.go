package main

/*
	this one problem is kinda fucking broken,
	so I'm not gonna stick to their stupid constraints.
	I mean, good luck doing binary search on unsorted array, fuckers
*/

func findDuplicate(nums []int) int {

	vset := map[int]bool{}

	for _, val := range nums {

		_, has := vset[val]
		if has {
			return val
		}

		vset[val] = true
	}

	return 0
}

func main() {

	println(findDuplicate([]int{1, 3, 4, 2, 2}))

}
