package main

import "fmt"

/*翻转含有 `中文、数字、英文字母` 等任意字符的字符串

示例

```go
输入："he師l發lo,世。+-*界，6"
输出："6，界*-+。世,ol發l師eh"
```*/

func reverseString(s string) string {
	// 将字符串转换为 rune 类型的切片，并对该切片翻转
	res := reverse([]int32(s))
	// 再把 rune 类型的切片转为 string
	return string(res)
}

func reverse(s []int32) []rune {
	// 左右指针，对切片依次翻转
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	s := "he師l發lo,世。+-*界，6"
	res := reverseString(s)
	fmt.Println(res)
}
