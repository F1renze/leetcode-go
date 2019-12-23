package q55

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	sum := 0
	for i := range nums {
		if i > sum {
			return false
		}
		// i + nums[i] 为当前点能达到的最远距离, 不断贪心
		sum = max(sum, i + nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 逆推, 更快
func canJump2(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	lens := len(nums)
	lp := lens - 1
	for i := lens - 1;i >= 0;i-- {
		if i + nums[i] >= lp {
			lp = i
		}
	}

	return lp == 0
}

