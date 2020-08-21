package main

import "fmt"

// 每个状态的下个状态都有 8 种, 即每个拨轮都可以有上下两种状态, 可以看成八叉树
func openLock(deadends []string, target string) int {
	// 记录需要跳过的死亡密码
	deadMap := make(map[string]bool)
	for i, deadEndsLen := 0, len(deadends); i < deadEndsLen; i++ {
		deadMap[deadends[i]] = true
	}

	// 记录已经穷举过的密码, 防止走回头路
	visited := make(map[string]bool)

	queue := make([]string, 0)

	// 从起点开始广度优先搜索
	step := 0
	queue = append(queue, "0000")
	visited["0000"] = true

	for len(queue) > 0 {
		queueLen := len(queue)
		// 将当前队列中的所有节点向周围扩散
		for i := 0; i < queueLen; i++ {
			// 过滤掉不符合的
			if deadMap[queue[i]] {
				continue
			}
			// 判断是否到达了终点
			if queue[i] == target {
				return step
			}
			// 将一个节点的未遍历相邻节点加入到队列
			for j := 0; j < 4; j++ {
				up := plusOne(queue[i], j)
				if !visited[up] {
					// 还没穷举过
					queue= append(queue, up)
					visited[up] = true
				}

				down := minusOne(queue[i], j)
				if !visited[down] {
					// 还没穷举过
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		queue = queue[queueLen:]
		// 增加步数
		step++
	}
	// 如果穷举完还没找到目标密码, 那就是找不到了
	return -1
}

// 将 s[j] + 1
func plusOne(s string, j int) string {
	res := []byte(s)
	if res[j] == '9' {
		res[j] = '0'
	} else {
		res[j]++
	}
	return string(res)
}

// s[j] - 1
func minusOne(s string, j int) string {
	res := []byte(s)
	if res[j] == '0' {
		res[j] = '9'
	} else {
		res[j]--
	}
	return string(res)
}

func main() {
	deadends := []string{"8887","8889","8878","8898","8788","8988","7888","9888"}
	target := "8888"
	res := openLock(deadends, target)
	fmt.Println(res)
}
