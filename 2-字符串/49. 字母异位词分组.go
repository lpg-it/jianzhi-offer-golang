package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	// 创建二维切片
	res := make([][]string, n)
	count := 0

	m := make(map[string]int, n)
	for _, v := range strs {
		s := strings.Split(v, "")
		sort.Strings(s)
		sortStr := strings.Join(s, "")

		if value, ok := m[sortStr]; ok {
			// 已存在
			res[value] = append(res[value], v)
		} else {
			// 不存在
			m[sortStr] = count
			res[count] = append(res[count], v)
			count += 1
		}
	}
	return res[:len(m)]
}

func groupAnagrams2(strs []string) [][]string {
	mWord := map[int32]int32{
		'a': 2,
		'b': 3,
		'c': 5,
		'd': 7,
		'e': 11,
		'f': 13,
		'g': 17,
		'h': 19,
		'i': 23,
		'j': 29,
		'k': 31,
		'l': 37,
		'm': 41,
		'n': 43,
		'o': 47,
		'p': 53,
		'q': 59,
		'r': 61,
		's': 67,
		't': 71,
		'u': 73,
		'v': 79,
		'w': 83,
		'x': 89,
		'y': 97,
		'z': 101,
	}
	n := len(strs)
	// 创建二维切片
	res := make([][]string, n)
	var count int32
	count = 0
	m := make(map[int32]int32, n)
	for _, v := range strs {
		var sumTemp int32
		sumTemp = 1
		for _, value := range v {
			sumTemp *= mWord[value]
		}

		if value, ok := m[sumTemp]; ok {
			// 已存在
			res[value] = append(res[value], v)
		} else {
			// 不存在
			m[sumTemp] = count
			res[count] = append(res[count], v)
			count += 1
		}
	}
	return res[:len(m)]
}

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	res := groupAnagrams2(a)

	fmt.Println(res)
}
