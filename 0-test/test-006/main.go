package main

import (
	"fmt"
)

/* 位操作 */
func main() {
	// 利用或操作 | 和空格将英文字符转换为小写
	fmt.Printf("%c\n", 'A' | ' ')
	// 利用与操作 & 和下划线将英文字符转换为大写
	fmt.Printf("%c\n", 'b' & '_')
	// 利用异或操作 ^ 和空格进行英文字符大小写互换
	fmt.Printf("%c\n", 'a' ^ ' ')
	fmt.Printf("%c\n", 'A' ^ ' ')
}
