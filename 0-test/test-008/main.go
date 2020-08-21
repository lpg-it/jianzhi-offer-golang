package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
/*	s := "hello"
	fmt.Println(&s)
	sli := []byte(s)
	fmt.Printf("%p", sli)*/
	s := "aaa"
	ssh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Println(b)


}
