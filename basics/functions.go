package main

import "fmt"

func functions() {
	sum := add(2, 3)
	fmt.Println(sum)

	greet := func() {
		fmt.Println("anonymous function")
	}

	greet()

	result := applyOperation(2, 4, add)
	fmt.Println(result)

	multiplyBy2 := createMultiplier(2)
	fmt.Println(multiplyBy2(6))
}

func add(x, y int) int {
	return x + y
}

// function that takes function as an argument
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(i int) int { return i * factor }
}
