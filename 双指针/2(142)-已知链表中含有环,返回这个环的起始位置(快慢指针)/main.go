package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	cycleFlag := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			cycleFlag = true
			break
		}
	}

	if !cycleFlag {
		// 不是环
		return nil
	}

	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {

}
