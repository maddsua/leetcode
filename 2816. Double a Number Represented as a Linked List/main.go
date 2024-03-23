package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {

	var number string
	for it := head; it != nil; it = it.Next {
		number += strconv.Itoa(it.Val)
	}

	var doubled string
	var carry int

	for i := len(number) - 1; i >= 0; i-- {

		posInt, _ := strconv.Atoi(string(number[i]))
		next := (posInt * 2) + carry

		if next < 10 {
			doubled += strconv.Itoa(next)
			carry = 0
			continue
		} else {
			carry = next / 10
			doubled += strconv.Itoa(next % 10)
		}
	}

	if carry > 0 {
		doubled += strconv.Itoa(carry)
	}

	var result *ListNode

	for _, character := range doubled {
		vAsInt, _ := strconv.Atoi(string(character))
		result = &ListNode{
			Val:  vAsInt,
			Next: result,
		}
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

func intToLL(number int64) *ListNode {

	numberAsString := reverseString(strconv.FormatInt(number, 10))

	var tempHead *ListNode

	for _, character := range numberAsString {
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

	number := intToLL(100500)
	printLLNumber(number)

	doubled := doubleIt(number)
	printLLNumber(doubled)

}
