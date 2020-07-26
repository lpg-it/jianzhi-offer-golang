package main

import "fmt"

func firstUniqChar(s string) byte {
	m := make(map[int32]bool, len(s))
	bs := make([]byte, 0, len(s))
	for _, value := range s {
		if _, ok := m[value]; !ok {
			// 没有
			m[value] = true
			bs = append(bs, byte(value))
		} else {
			// 已存在
			m[value] = false
		}
	}
	for _, value := range bs {
		if m[int32(value)] {
			return value
		}
	}
	return ' '
}

func main() {
	str := "loveleetcode"
	fmt.Println(string(firstUniqChar(str)))
}
