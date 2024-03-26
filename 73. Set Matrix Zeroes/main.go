package main

import "fmt"

type coordinate struct {
	row int
	col int
}

func setZeroes(matrix [][]int) {

	zeros := make([]coordinate, 0)

	for x, row := range matrix {

		for y, v := range row {

			if v == 0 {
				zeros = append(zeros, coordinate{
					row: x,
					col: y,
				})
			}
		}
	}

	for _, coord := range zeros {

		//	set y'es
		for col := range matrix[coord.row] {
			matrix[coord.row][col] = 0
		}

		//	set x'es
		for row := range matrix {
			matrix[row][coord.col] = 0
		}
	}
}

func main() {

	data := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(data)

	for _, m := range data {
		for _, n := range m {
			fmt.Printf("%d ", n)
		}
		fmt.Printf("\n")
	}
}
