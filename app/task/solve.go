package task

import (
	"sort"
)

// My version func for task Solve
func Solve2(arr []int) []int {
	result := make([]int, 0, len(arr))

	frequency := make(map[int]int)
	for _, value := range arr {
		frequency[value]++
	}

	var sorted []struct{ Key, Value int }

	for k, v := range frequency {
		sorted = append(sorted, struct{ Key, Value int }{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	for _, value := range sorted {
		for i := 0; i < value.Value; i++ {
			result = append(result, value.Key)
		}
	}

	return result
}

// Version ChatGPT
func Solve1(arr []int) []int {
	frequency := make(map[int]int)
	for _, value := range arr {
		frequency[value]++
	}

	var sorted []struct{ Key, Value int }

	for k, v := range frequency {
		sorted = append(sorted, struct{ Key, Value int }{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	result := make([]int, 0, len(arr))

	for _, item := range sorted {
		for i := 0; i < item.Value; i++ {
			result = append(result, item.Key)
		}
	}

	return result
}

// Version CodeWars
func Solve(arr []int) []int {
	counter := map[int]int{}
	for _, n := range arr {
		counter[n]++
	}
	sort.SliceStable(arr, func(i, j int) bool {
		m, n := arr[i], arr[j]
		a, b := counter[m], counter[n]
		return a > b || (a == b && m < n)
	})
	return arr
}
