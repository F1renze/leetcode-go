package q455

func findContentChildren(g []int, s []int) int {
	if len(g) < 1 || len(s) < 1 {
		return 0
	}
	// quick sort faster than sort.Slice
	quickSort(g, 0, len(g)-1)
	quickSort(s, 0, len(s)-1)

	var c, b int
	for c < len(g) && b < len(s) {
		if s[b] >= g[c] {
			c++
		}
		b++
	}

	return c
}

func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}

	l, r, t := low, high, nums[low]
	for l < r {
		for l < r && t <= nums[r] {
			r--
		}
		nums[l] = nums[r]
		for l < r && t >= nums[l] {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = t

	quickSort(nums, low, l-1)
	quickSort(nums, l+1, high)
}
