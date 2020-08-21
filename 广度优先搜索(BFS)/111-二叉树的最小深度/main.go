package main

import "math"

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

// 广度优先
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	// root 本身就是一层, depth 初始化为 1
	depth := 1

	for len(queue) > 0 {
		queueLen := len(queue)
		// 将当前队列中的所有节点向四周扩散
		for i := 0; i < queueLen; i++ {
			// 判断是否到达终点
			if queue[i].Left == nil && queue[i].Right == nil {
				return depth
			}
			// 将当前节点的相邻节点加入队列
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// 已经计算过的节点无需重复计算
		queue = queue[queueLen:]
		// 在这里增加步数
		depth++
	}
	return depth
}

// 深度优先
var min int
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 只有根节点的深度为 1
	depth := 1

	min = math.MaxInt32

	dfs(root, depth)
	return min
}

func dfs(node *TreeNode, depth int) {
	// 判断结束条件
	if node.Left == nil && node.Right == nil {
		if depth < min {
			min = depth
		}
		return
	}

	if node.Left != nil {
		dfs(node.Left, depth+1)
	}
	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}

func main() {

}
