package task

import (
	"strconv"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func splitStep(s string, step int) []string {
	var result []string

	for i := 0; i < len(s); i += step {
		end := i + step
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}

	return result
}

// O(n) == O(len(n))
func Revrot(s string, n int) string {
	if s == "" || n == 0 || len(s) < n {
		return ""
	}

	sliceNumber := splitStep(s, n)
	var resultList []string

	for _, l := range sliceNumber {

		if len(l) < n {
			continue
		}

		sum := 0

		for _, t := range l {
			num, _ := strconv.Atoi(string(t))
			sum += num
		}
		if sum%2 == 0 {
			resultList = append(resultList, reverseString(l))
		} else {
			resultList = append(resultList, l[1:]+l[:1])
		}
	}

	return strings.Join(resultList, "")
}
