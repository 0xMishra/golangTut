package main

import "fmt"

func switchCasesTut() {
	// there is no break in switch statements in go like other languages
	// we can use fallthrough to check for other cases even if a case is satisfied
	switch age := 25; age {
	case 25:
		fmt.Println("you are 25")
	case 30:
		fmt.Println("You are 30")
	default:
		fmt.Println("You are neither 25 nor 30")
	}

	day := "monday"
	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("It is a week day")
	case "saturday", "sunday":
		fmt.Println("Its a weekend")
	default:
		fmt.Println("Its an invalid day")
	}

	number := 18
	switch {
	case number >= 18:
		fmt.Println("You are an adult")
	case number < 18:
		fmt.Println("You are not an adult")
	default:
		fmt.Println("invalid age")
	}

	checkType(23)
}

func checkType(x any) {
	// we can't use fallthrough here in type switch case
	switch x.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		fmt.Println("x is a number")
	case float32, float64:
		fmt.Println("x is a float")
	case string:
		fmt.Println("x is a string")
	case bool:
		fmt.Println("x is a boolean")
	default:
		fmt.Println("x is an unknown type")
	}
}
