package task

import (
	"math"
	"strconv"
)

func ReverseInteger(x int) int {
	sI := []byte(strconv.Itoa(x))

	i, j := 0, len(sI)-1
	for i < j {
		if string(sI[i]) == "-" {
			i++
		}
		// Является ли ноль последним?
		if string(sI[j]) == "0" && len(sI) == j+1 {
			sI = sI[:j]
			j--
			continue
		}
		sI[i], sI[j] = sI[j], sI[i]
		i++
		j--
	}

	r, _ := strconv.Atoi(string(sI))

	y := int(math.Pow(2, 32))/2 - 1
	if r > y || (r*-1) > y {
		return 0
	}

	return r
}

// LeetCode
func ReverseInteger1(x int) int {
	reversed := 0

	for x != 0 {
		digit := x % 10                // get last digit
		reversed = reversed*10 + digit // add as new digit
		x /= 10                        // update x
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	return reversed
}

func ReverseInteger2(x int) int {
	reversed := 0

	for x != 0 {
		digit := x % 10
		x /= 10
		if reversed > (1<<31-1)/10 || reversed < -(1<<31)/10 {
			return 0
		}

		reversed = reversed*10 + digit
	}

	return reversed
}
