package main

type Node struct {
	Val int
	Next *Node
}

func hasCycle(head *Node) bool {
	var (
		fast *Node
		slow *Node
	)
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			// 是环
			return true
		}
	}
	return false
}

func main() {

}
