package main

import "fmt"

func main() {
	/*	// 利用 或操作 | 和 空格 将英文字符转换为小写
		a := 'a' | ' '
		b := 'B' | ' '
		fmt.Printf("%c, %c", a, b)*/

	/*	// 利用与操作 `&` 和下划线将英文字符转换为大写
		A := 'a' & '_'
		B := 'B' & '_'
		fmt.Printf("%c, %c", A, B)*/

	/*	// 利用异或操作 `^` 和空格进行英文字符大小写互换
		A := 'a' ^ ' '
		b := 'B' ^ ' '
		fmt.Printf("%c, %c", A, b)  // A, b*/

	/*	// 判断两个数是否异号
		x, y := 1, -2
		f := x^y < 0
		fmt.Println(f) // true

		a, b := 1, 2
		f = a^b < 0
		fmt.Println(f)  // false*/

	// 不用临时变量交换两个数
	a, b := 1, 2
	a ^= b  // 3, 2
	b ^= a  // 3, 1
	a ^= b  // 2, 1
	fmt.Println(a, b)
}
