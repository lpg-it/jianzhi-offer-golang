## 题目描述

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

### 示例1

```go
输入：s = "We are happy."
输出："We%20are%20happy."
```

## Golang实现

### 方法一

增加一个额外空间，一个一个对比，碰到空格添加 %20，否则添加原字符。

```go
package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	str := ""
	for _, value := range s {
		if value == ' ' {
			str += "%20"
		} else {
			str += string(value)
		}
	}
	return  str
}

func main() {
	s := "We are happy."
	newS := replaceSpace(s)
	fmt.Printf(newS)
}
```

执行用时 : 0 ms, 在所有 Go 提交中击败了 100% 的用户

内存消耗 : 3.3 MB, 在所有 Go 提交中击败了 100% 的用户

### 方法二

直接先按照空格分割成切片，再按照 %20 拼接成字符串即可，一行解决。

```go
package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.Join(strings.Split(s, " "), "%20")
}

func main() {
	s := "We are happy."
	newS := replaceSpace(s)
	fmt.Printf(newS)
}
```

执行用时 : 0 ms, 在所有 Go 提交中击败了 100% 的用户

内存消耗 : 2 MB, 在所有 Go 提交中击败了 100% 的用户

最后附上本人的LeetCode题解链接：[https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/solution/ti-huan-kong-ge-liang-chong-jie-fa-by-lpgit/](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/solution/ti-huan-kong-ge-liang-chong-jie-fa-by-lpgit/)

## 李培冠博客

[lpgit.com](https://lpgit.com)
