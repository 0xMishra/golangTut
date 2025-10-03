package main

import "fmt"

func pointers() {
	a := 10
	ptr := &a

	fmt.Println(a)
	fmt.Println(ptr)

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
