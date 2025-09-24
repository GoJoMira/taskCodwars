package task

import (
	"math/big"
	"strconv"
	"strings"
)

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/559/

// Пришлось использовать math/big

// Мое решение. Runtime < 1ms. Memory 4300 KB
func PlusOne(digits []int) []int {
	var result []int

	var y []string

	for _, i := range digits {
		y = append(y, strconv.Itoa(i))
	}

	x := strings.Join(y, "")

	n := new(big.Int)
	u, _ := n.SetString(x, 10)
	u.Add(u, big.NewInt(1))

	u2 := strings.Split(n.String(), "")

	for _, i := range u2 {
		r, _ := strconv.Atoi(i)
		result = append(result, r)
	}

	return result
}

// Другие решения из LeetCode
