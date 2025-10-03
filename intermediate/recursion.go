package main

import "fmt"

func recursion() {
	fmt.Println("factorial of 5:", factorial(5))
	fmt.Println("sum of digits in the number 1289:", sumOfDigits(1289))
}

func factorial(n int) int {
	// base case
	if n == 0 {
		return 1
	}

	// factorial case
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// base case
	if n < 10 && n > 0 {
		return n
	}

	// factorial case
	return n%10 + sumOfDigits(n/10)
}
