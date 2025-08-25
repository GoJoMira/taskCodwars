package task

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/674/

func Intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) == 1 && len(nums2) == 1 {

		if nums1[0] == nums2[0] {
			return nums1
		}

		return []int{}
	}

	var first, last []int
	var result []int

	if len(nums1) > len(nums2) {
		first, last = nums1, nums2
	} else {
		first, last = nums2, nums1
	}

	indexFirstNumber := make(map[int]int)
	indexlastNumber := make(map[int]int)

	for i, v := range first {
	OuterLoop:
		for i2, v2 := range last {

			if v == v2 {
				_, ok := indexlastNumber[i2]
				_, ok1 := indexFirstNumber[i]
				if !ok && !ok1 {
					indexFirstNumber[i] = 0
					indexlastNumber[i2] = 0
					result = append(result, v)
					break OuterLoop
				} else {
					continue
				}
			}

		}
	}

	return result
}

// Решение от LeetCode
