package main

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	temp := &ListNode{Next: head}

	for it := temp; it.Next != nil && it.Next.Next != nil; {
		nextNext := it.Next.Next
		next := it.Next
		next.Next = nextNext.Next
		nextNext.Next = next
		it.Next = nextNext
		it = next
	}

	return temp.Next
}

func arrayToNodelist(data []int) *ListNode {

	var temp *ListNode = nil

	for i := len(data) - 1; i >= 0; i-- {
		temp = &ListNode{Val: data[i], Next: temp}
	}

	return temp
}

func printNodeList(list *ListNode) {
	temp := ""
	for it := list; it != nil; it = it.Next {
		temp += strconv.Itoa(it.Val) + ", "
	}
	println(temp)
}

func main() {

	list := arrayToNodelist([]int{1, 2, 3, 4})
	res := swapPairs(list)
	printNodeList(res)
}
