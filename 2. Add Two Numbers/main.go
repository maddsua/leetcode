package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var carry int8
	var sum []int8

	for it1, it2 := l1, l2; it1 != nil || it2 != nil; {

		var val1 int8
		if it1 != nil {
			val1 = int8(it1.Val)
			it1 = it1.Next
		}

		var val2 int8
		if it2 != nil {
			val2 = int8(it2.Val)
			it2 = it2.Next
		}

		tempSum := val1 + val2 + carry

		if tempSum < 10 {
			carry = 0
			sum = append(sum, tempSum)
		} else {
			carry = 1
			sum = append(sum, tempSum-10)
		}
	}

	if carry > 0 {
		sum = append(sum, carry)
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
	numB := numberToLL(1007)
	printLLNumber(numA)
	printLLNumber(numB)

	sum := addTwoNumbers(numA, numB)
	printLLNumber(sum)

}
