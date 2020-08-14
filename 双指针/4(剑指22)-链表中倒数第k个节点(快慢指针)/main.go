package main

type ListNode struct {
	Val int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k && fast != nil; i++{
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	
}
