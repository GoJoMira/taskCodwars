package task

//10600 KB Memory
func Rotate(nums []int, k int) {

	if len(nums) == 1 {
		return
	}

	if len(nums) == k {
		return
	}

	result := make([]int, len(nums))

	for i, v := range nums {

		sumI := i + k
		numsL := len(nums)

		if sumI >= numsL {
			for sumI >= numsL {
				sumI -= numsL
			}
			result[sumI] = v
		} else {
			result[sumI] = v
		}
	}
	copy(nums, result)
}

// Code LeetCode
// 8800 KB
func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func Rotate1(nums []int, k int) {
	k %= len(nums)

	Reverse(nums)
	Reverse(nums[:k])
	Reverse(nums[k:])
}
