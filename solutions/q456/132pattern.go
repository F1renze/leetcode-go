package q456


func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}

	var (
		s []int
	)
	m := make([]int, n)
	m[0] = nums[0]
	for i := 1; i < n; i++ {
		m[i] = min(m[i-1], nums[i])
	}

	for i := n-1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= m[i] {
			s = s[:len(s)-1]
		}
		if len(s) > 0 && s[len(s)-1] < nums[i] {
			return true
		}
		s = append(s, nums[i])
	}

	return false
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
