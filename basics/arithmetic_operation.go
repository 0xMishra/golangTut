package basics

import "fmt"

func arithOp() {
	a, b := 10, 3

	var result int

	result = a + b
	fmt.Println("Additon: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b      // returns an int
	const pi = 22 / 7.0 // returns a float
	fmt.Println("Division: ", result, pi)

	result = a % b
	fmt.Println("Modulus: ", result)

	var maxInt int64 = 987654321987654321
	maxInt++

	fmt.Println(maxInt)
}
