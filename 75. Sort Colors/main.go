package main

func sortColors(nums []int) {

	itemsMap := map[int]int{}

	for _, v := range nums {

		counter, has := itemsMap[v]
		if !has {
			itemsMap[v] = 1
			continue
		}

		itemsMap[v] = counter + 1
	}

	var gi int

	for _, v := range []int{0, 1, 2} {

		count, has := itemsMap[v]
		if !has {
			continue
		}

		for li := 0; li < count; li++ {
			nums[gi] = v
			gi++
		}
	}
}

func printColors(nums []int) {
	for _, v := range nums {
		print(v)
	}
	print("\n")
}

func main() {

	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	printColors(colors)
}
