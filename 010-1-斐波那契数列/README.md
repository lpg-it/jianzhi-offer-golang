## 题目描述

写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

```go
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
```

斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

### 示例1

```go
输入：n = 2
输出：1
```

### 示例2

```go
输入：n = 5
输出：5
```

## Golang实现

注意看下面的规律，下一个 x 总是等于上一个 y，下一个 y 总是等于上一个 x+y

```go
0 1 1 2 3 5 8 13
x y
  x y
    x y
      ...
```

所以，我们可以在 Go 中直接赋值：`x, y = y, x+y`，这样一来，外面套个循环依次计算就完事了。

```go
package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y%1000000007, (x+y)%1000000007
	}
	return x
}

func main() {
	n := 40
	fmt.Println(fib(n))
}
```

执行用时 : 0 ms, 在所有 Go 提交中击败了 100% 的用户

内存消耗 : 1.9 MB, 在所有 Go 提交中击败了 100% 的用户

## 用 select 实现

题外话，我们也可以用 select 去实现一个依次打印斐波那契数列的程序。

```go
package main

import (
	"fmt"
	"runtime"
)

func fib(ch <- chan int, quit <- chan bool) int {
	for  {
		select {
		case num := <- ch:
			fmt.Println(num)
		case <- quit:
			runtime.Goexit()
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go fib(ch, quit)

	x, y := 0, 1
	for i := 0; i < 30; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
```

最后附上本人的LeetCode题解链接：[https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/solution/fei-bo-na-qi-shu-lie-shuang-bai-yong-select-shi-xi/](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/solution/fei-bo-na-qi-shu-lie-shuang-bai-yong-select-shi-xi/)

## 李培冠博客

[lpgit.com](https://lpgit.com)
