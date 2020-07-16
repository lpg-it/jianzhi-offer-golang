package main

import (
	"container/list"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*func reversePrint(head *ListNode) []int {
	var s []int
	// 顺序存储数据到切片中
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	// 反转切片数据的位置
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}*/

// 用栈
func reversePrint(head *ListNode) []int {
	var s []int
	stack := list.New()
	// 链表中数据顺序入栈
	for head != nil {
		stack.PushBack(head.Val)
		head = head.Next
	}
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		s = append(s, e.Value.(int))
	}
	return s
}

func main() {
	var head ListNode
	var node ListNode
	head.Val = 5
	node.Val = 8
	head.Next = &node

	s := reversePrint(&head)
	fmt.Println(s)
}
