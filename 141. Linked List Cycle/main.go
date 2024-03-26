package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	nodeSet := make(map[*ListNode]bool, 0)

	for it := head; it != nil; it = it.Next {

		_, has := nodeSet[it]
		if has {
			return true
		}

		nodeSet[it] = true
	}

	return false
}

func main() {

}
