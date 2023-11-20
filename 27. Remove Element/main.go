package main

func removeElement(nums []int, val int) int {

	temp := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			temp = append(temp, nums[i])
		}
	}

	for i := 0; i < len(temp); i++ {
		nums[i] = temp[i]
	}

	return len(temp)
}

func main() {

	data := []int{0,1,2,2,3,0,4,2}

	notEqual := removeElement(data, 2)

	println("Not equal:", notEqual)
	println("elements:")

	for i := 0; i < len(data); i++ {
		print(" ", data[i])
	}

}
