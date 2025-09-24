package task

func TwoSum(nums []int, target int) []int {
	index := make([]int, 2)
	mapDifference := make(map[int]int)

	for i, v := range nums {
		n := target - v
		if _, ok := mapDifference[n]; !ok {
			mapDifference[v] = i
		} else {
			index[0] = mapDifference[n]
			index[1] = i
		}
	}

	return index
}

// LeetCode
// Не могу понять. В данной функции есть вероятность
// что вернется nil, а это совем не []int, почему компилятор не ругается?

// Крч... []int map[int]int можно возвращать nil
// int, bool, struct{} тут уже nil не вернуть.
func TwoSum1(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}
