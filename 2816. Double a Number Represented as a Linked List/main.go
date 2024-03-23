package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(head *ListNode, value int) *ListNode {
	vAsInt, _ := strconv.Atoi(strconv.Itoa(value))
	return &ListNode{
		Val:  vAsInt,
		Next: head,
	}
}

func doubleIt(head *ListNode) *ListNode {

	var number string
	for it := head; it != nil; it = it.Next {
		number += strconv.Itoa(it.Val)
	}

	var result *ListNode
	var carry int

	for i := len(number) - 1; i >= 0; i-- {

		posInt, _ := strconv.Atoi(string(number[i]))
		next := (posInt * 2) + carry

		if next < 10 {
			result = addNode(result, next)
			carry = 0
			continue
		} else {
			carry = next / 10
			result = addNode(result, next%10)
		}
	}

	if carry > 0 {
		result = addNode(result, carry)
	}

	return result
}

func reverseString(s string) string {

	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func numberToLL(number string) *ListNode {

	temp := reverseString(number)

	var tempHead *ListNode

	for _, character := range temp {
		vAsInt, _ := strconv.Atoi(string(character))
		tempHead = &ListNode{
			Val:  vAsInt,
			Next: tempHead,
		}
	}

	return tempHead
}

func printLLNumber(head *ListNode) {
	for it := head; it != nil; it = it.Next {
		fmt.Printf("%d", it.Val)
	}
	fmt.Printf("\n")
}

func main() {

	number := numberToLL("100500")
	printLLNumber(number)

	doubled := doubleIt(number)
	printLLNumber(doubled)

}
