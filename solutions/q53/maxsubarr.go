package q53

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	r, sum := nums[0], 0

	for i := range nums {
		if sum > 0 {
			// faster
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > r {
			r = sum
		}
	}
	return r
}
