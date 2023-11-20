package main

func plusOne(digits []int) []int {

	for idx := len(digits) - 1; idx >= 0; idx-- {

		if digits[idx] < 9 {
			digits[idx]++
			return digits
		}

		digits[idx] = 0
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {

	res := plusOne([]int{1, 9, 9, 9})

	for _, v := range(res) {
		print(v)
	}

}