package task

import (
	"math"
)

// My solution
func Century(year int) int {
	var a int = 0
	num := float64(year) * 0.01
	result := math.Mod(num, 1)
	if result > 0 {
		a = int(num) + 1
	} else {
		a = int(num)
	}
	return a
}

// Solution with codewars

// func century(year int) int {
// 	return (year + 99) / 100
// }

// func century(y int) int {
// 	return int(math.Ceil(float64(y) / 100))
// }

// ChatGPT also decided the same.
// func century(year int) int {
// 	// Finish this :)
// 	if year%100 != 0 {
// 		year += 100
// 	}
// 	return year / 100
// }
