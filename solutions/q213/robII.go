package q213

func rob(nums []int) int {
	size := len(nums)
	if size < 1 {
		return 0
	}
	if size < 2 {
		return nums[0]
	}

	a := robHelper(nums[:size-1])
	b := robHelper(nums[1:])

	if a > b {
		return a
	} else {
		return b
	}
}

func robHelper(nums []int) int {
	var prev, cur int

	for i := range nums {
		if (prev + nums[i]) > cur {
			prev, cur = cur, prev + nums[i]
		} else {
			prev = cur
		}
	}
	return cur
}