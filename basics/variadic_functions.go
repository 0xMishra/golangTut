package basics

import "fmt"

func variadicFunc() {
	// ...elipsis
	seq, total := sum(1, 2, 3, 4)
	fmt.Println("Sum is :", total, "sequence is", seq)

	// passing the slice as variadic paramters
	numbers := []int{2, 30, 40, 50, 60}
	sequence, total2 := sum(2, numbers...)
	fmt.Println("Sum is :", total2, "sequence is", sequence)
}

// variadic paramters must be the last paramters in function signature
func sum(sequence int, nums ...int) (int, int) {
	total := 0

	for _, v := range nums {
		total += v
	}

	return sequence, total
}
