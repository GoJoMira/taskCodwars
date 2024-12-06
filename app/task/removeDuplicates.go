package task

import "sort"

// My Version
func RemoveDuplicates1(nums []int) int {
	uniqueMapNumbers := map[int]int{}
	for i := 0; i < len(nums); i++ {
		uniqueMapNumbers[nums[i]]++
	}

	var sorted []struct{ Key, Value int }

	for k, v := range uniqueMapNumbers {
		sorted = append(sorted, struct{ Key, Value int }{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Key < sorted[j].Key
	})

	keyUniqueMapNumbers := []int{}
	for _, value := range sorted {
		keyUniqueMapNumbers = append(keyUniqueMapNumbers, value.Key)
	}

	copy(nums, keyUniqueMapNumbers)

	return len(uniqueMapNumbers)
}

// Version 2
func RemoveDuplicates(nums []int) int {
	uniqueMap := make(map[int]struct{})
	for _, num := range nums {
		uniqueMap[num] = struct{}{}
	}

	keys := make([]int, 0, len(uniqueMap))
	for k := range uniqueMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	copy(nums, keys)

	return len(keys)
}
