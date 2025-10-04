package basics

import (
	"errors"
	"fmt"
)

func multiReturnVal() {
	q, r := divide(10, 3)
	fmt.Println(q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a>b", nil
	} else if b > a {
		return "b>a", nil
	}
	return "", errors.New("unable to compare which is greater")
}
