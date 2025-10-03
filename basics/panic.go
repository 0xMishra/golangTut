package main

import "fmt"

func panicTut() {
	panicProcess(10)

	panicProcess(-10)
}

func panicProcess(input int) {
	defer fmt.Println("defered statement 1")
	defer fmt.Println("defered statement 2")
	if input < 0 {
		fmt.Println("Before panic")
		// this statement will be printed at last after all the defer statements have been executed
		panic("Input should be a positive number")
	}
	fmt.Println("Processing input:", input)
}
