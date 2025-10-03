package main

import "fmt"

func ifElse() {
	if age := 25; age >= 18 {
		fmt.Println("You are an adult")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	// nested if else
	num := 15
	if num%2 == 0 {
		if num%3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3")
		} else {
			fmt.Println("Number is divisible by only 2 not 3")
		}
	}
}
