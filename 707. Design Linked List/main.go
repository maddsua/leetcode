package main

import (
	"strconv"
	"strings"
)

type LinkedListItem struct {
	Value int
	Next  *LinkedListItem
}

type MyLinkedList struct {
	Head *LinkedListItem
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {

	if index < 0 || index >= this.Size {
		return -1
	}

	for it, i := this.Head, 0; it != nil; it, i = it.Next, i+1 {
		if i == index {
			return it.Value
		}
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {

	next := LinkedListItem{
		Value: val,
		Next:  this.Head,
	}

	this.Head = &next
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {

	var tail *LinkedListItem
	for it := this.Head; it != nil; it = it.Next {
		tail = it
	}

	next := LinkedListItem{
		Value: val,
		Next:  nil,
	}

	if tail == nil {
		this.Head = &next
	} else {
		tail.Next = &next
	}

	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index > this.Size {
		return
	}

	if index == this.Size {
		this.AddAtTail(val)
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	for it, i := this.Head, 0; it != nil; it, i = it.Next, i+1 {

		if i != index-1 {
			continue
		}

		next := LinkedListItem{
			Value: val,
			Next:  it.Next,
		}

		it.Next = &next

		break
	}

	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	if index < 0 || index >= this.Size {
		return
	}

	if index == 0 {

		if this.Head == nil {
			return
		}

		this.Head = this.Head.Next
	}

	for it, i := this.Head, 0; it != nil; it, i = it.Next, i+1 {

		if i != index-1 {
			continue
		}

		var next *LinkedListItem

		if it.Next != nil {
			next = it.Next.Next
		}

		it.Next = next

		break
	}

	this.Size--
}

func printAll(head *LinkedListItem) {

	items := make([]string, 0)

	for it := head; it != nil; it = it.Next {
		items = append(items, strconv.Itoa(it.Value))
	}

	println(strings.Join(items, ","))
}

func main() {

	ll := Constructor()

	ll.AddAtHead(1)
	ll.AddAtTail(3)

	ll.AddAtIndex(1, 2)
	println(ll.Get(1))

	//printAll(ll.Head)

	ll.DeleteAtIndex(0)

	printAll(ll.Head)

	println(ll.Get(0))
}
