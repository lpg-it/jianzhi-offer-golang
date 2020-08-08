package main

import (
	"fmt"
	"io"
)

/* 牛客网做算法题时 golang 读取输入的方法 */
func main() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(a + b)
	}
}