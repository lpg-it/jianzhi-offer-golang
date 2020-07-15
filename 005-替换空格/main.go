package main

import (
	"fmt"
	"strings"
)

/*func replaceSpace(s string) string {
	str := ""
	for _, value := range s {
		if value == ' ' {
			str += "%20"
		} else {
			str += string(value)
		}
	}
	return  str
}*/

func replaceSpace(s string) string {
	return strings.Join(strings.Split(s, " "), "%20")
}

func main() {
	s := "We are happy."
	newS := replaceSpace(s)
	fmt.Printf(newS)
}
