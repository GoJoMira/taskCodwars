package task

import (
	"sort"
	"strings"
)

// Алгоритм до 3999
var nRim map[int]string = map[int]string{
	1:     "I",
	4:     "IV",
	5:     "V",
	9:     "IX",
	10:    "X",
	40:    "XL",
	50:    "L",
	90:    "XC",
	100:   "C",
	400:   "CD",
	500:   "D",
	900:   "CM",
	1000:  "M",
	10000: "0",
}

func getNRim(n int) string {
	var result string

	// TODO: Тут не должна быть генерация ключей
	// Данный цикл очень сильно замедляет функцию
	var keys []int
	for k := range nRim {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// var keys []int = []int{10000, 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	// sort.Ints(keys)

	if v := nRim[n]; v != "" {
		return v
	}

	for k := 0; k < len(keys)-1; k++ {
		if keys[k] < n && n < keys[k+1] {
			result += nRim[keys[k]]
			result += getNRim(n - keys[k])
		}
	}

	return result
}

func splitNumber(n int) []int {
	r := []int{}
	for d := 1; n/d > 0; d *= 10 {
		num := (n / d % 10) * d
		r = append([]int{num}, r...)
	}
	return r
}

func Solution(number int) string {
	var result string

	numbersList := splitNumber(number)

	for _, n := range numbersList {
		result += getNRim(n)
	}

	return result
}

// Решение с CodeWars
func Solution1(n int) string {
	if n <= 0 {
		return ""
	}
	var sb strings.Builder
	for i, v := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for n >= v {
			n -= v
			sb.WriteString([]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}[i])
		}
	}
	return sb.String()
}
