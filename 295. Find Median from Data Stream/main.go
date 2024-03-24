package main

import (
	"fmt"
	"sort"
)

type MedianFinder struct {
	data []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		data: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.data = append(this.data, num)
}

func (this *MedianFinder) FindMedian() float64 {

	sort.Ints(this.data)

	if len(this.data) == 0 {
		return 0
	}

	if len(this.data) == 1 {
		return float64(this.data[0])
	}

	if (len(this.data) % 2) != 0 {
		return float64(this.data[len(this.data)/2])
	}

	left := this.data[(len(this.data)/2)-1]
	right := this.data[(len(this.data) / 2)]

	return (float64(left) + float64(right)) / 2
}

func main() {

	mf := Constructor()

	mf.AddNum(6)
	//fmt.Printf("%.2f\n", mf.FindMedian())
	mf.AddNum(10)
	//fmt.Printf("%.2f\n", mf.FindMedian())
	mf.AddNum(2)

	fmt.Printf("%v\n", mf.data)

	fmt.Printf("%.2f\n", mf.FindMedian())
}
