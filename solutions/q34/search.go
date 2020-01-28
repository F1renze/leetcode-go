package q34

func searchRange(nums []int, target int) []int {
	l, r, m := 0, len(nums)-1, 0
	arr := []int{-1, -1}
	if r < 0 {
		return arr
	}

	for l < r {
		m = l + (r - l) / 2
		// 移动到左边界
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if nums[l] != target {
		return arr
	}
	arr[0], r = l, len(nums)

	for l < r {
		m = l + (r - l) / 2
		// 移动到右边界
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	arr[1] = r - 1

	return arr
}