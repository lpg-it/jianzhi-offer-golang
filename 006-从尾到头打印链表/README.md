## 题目描述

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

### 示例1

```go
输入：head = [1,3,2]
输出：[2,3,1]
```

## Golang实现

### 方法一

顺序取出链表中的数据存储到切片中, 然后返回切片的逆序即可.

```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
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
```

执行用时 : 0 ms, 在所有 Go 提交中击败了 100% 的用户

内存消耗 : 3.1 MB, 在所有 Go 提交中击败了 100% 的用户

### 方法二

用栈

```go
package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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
```

执行用时 : 4 ms, 在所有 Go 提交中击败了 60.58% 的用户

内存消耗 : 3.7 MB, 在所有 Go 提交中击败了 100% 的用户

最后附上本人的LeetCode题解链接：[https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/solution/cong-wei-dao-tou-da-yin-lian-biao-liang-chong-jie-/](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/solution/cong-wei-dao-tou-da-yin-lian-biao-liang-chong-jie-/)

## 李培冠博客

[lpgit.com](https://lpgit.com)
