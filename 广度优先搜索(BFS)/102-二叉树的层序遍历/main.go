package main

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		queueLen := len(queue)
		res = append(res, []int{})
		for i := 0; i < queueLen; i++ {
			res[len(res)-1] = append(res[len(res)-1], queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[queueLen:]
	}
	return res
}

func main() {

}
