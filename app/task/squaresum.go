package task

func SquareSum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}
