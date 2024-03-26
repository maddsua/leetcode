package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func llNumberToSlice(head *ListNode) []int8 {

	var nums []int8

	for it := head; it != nil; it = it.Next {
		nums = append(nums, int8(it.Val))
	}

	return nums
}

func selectLonger(a []int8, b []int8) ([]int8, []int8) {

	if len(a) >= len(b) {
		return a, b
	}

	return b, a
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	primary, secondary := selectLonger(llNumberToSlice(l1), llNumberToSlice(l2))

	var carry bool
	var sum []int8

	for idx, val := range primary {

		var sval int8
		if idx < len(secondary) {
			sval = secondary[idx]
		}

		temp := val + sval
		if carry {
			temp++
		}

		if temp < 10 {
			carry = false
			sum = append(sum, temp)
		} else {
			carry = true
			sum = append(sum, temp-10)
		}
	}

	if carry {
		sum = append(sum, 1)
	}

	var result *ListNode
	for idx := range sum {
		result = &ListNode{
			Next: result,
			Val:  int(sum[len(sum)-1-idx]),
		}
	}

	return result
}

func numberToLL(number int) *ListNode {

	var result *ListNode

	for _, character := range strconv.Itoa(number) {
		result = &ListNode{
			Val:  int(character) - 48,
			Next: result,
		}
	}

	return result
}

func printLLNumber(head *ListNode) {
	for it := head; it != nil; it = it.Next {
		fmt.Printf("%d", it.Val)
	}
	fmt.Printf("\n")
}

func main() {

	numA := numberToLL(123)
	numB := numberToLL(250)
	printLLNumber(numA)
	printLLNumber(numB)

	sum := addTwoNumbers(numA, numB)
	printLLNumber(sum)

}
