package main

import (
	"math/rand"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Size int
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	size := 0
	for it := head; it != nil; it = it.Next {
		size++
	}
	return Solution{Size: size, Head: head}
}

func (this *Solution) GetRandom() int {
	randomidx := rand.Intn(this.Size)

	for it, idx := this.Head, 0; it != nil; it, idx = it.Next, idx+1 {
		if idx == randomidx {
			return it.Val
		}
	}

	return this.Head.Val
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

	list := arrayToNodelist([]int{1, 3, 2, 2, 3})
	obj := Constructor(list)
	param_1 := obj.GetRandom()

	println(param_1)

}
