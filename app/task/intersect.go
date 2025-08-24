package task

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/674/

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) == 1 && len(nums2) == 1 {
		return nums1
	}

	var first, last []int
	var result []int

	if len(nums1) > len(nums2) {
		first, last = nums1, nums2
	} else {
		first, last = nums2, nums1
	}

	for _, v := range last {
		for _, i := range first {
			if v == i {
				result = append(result, v)
			}
		}
	}

	var fainal_result []int

	sortNums(result)

	for i, j := 0, 1; j < len(result); i, j = i+1, j+1 {
		if result[i] == result[j] {
			fainal_result = append(fainal_result, result[i])
		}
	}

	return fainal_result
}

func sortNums(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] == nums[j] {
			return nums[i] > nums[j]
		}
		return nums[i] < nums[j]
	})
}
