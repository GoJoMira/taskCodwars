package task

func SquareSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}
