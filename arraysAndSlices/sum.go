package arraysandslices

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, arr := range numbersToSum {
		sum = append(sum, Sum(arr))
	}

	return sum
}

func SumAllTails(numsToSum ...[]int) []int {
	var sum []int

	for _, arr := range numsToSum {
		if len(arr) == 0 {
			sum = append(sum, 0)
		} else {
			arr = arr[1:]
			sum = append(sum, Sum(arr))
		}
	}

	return sum
}
