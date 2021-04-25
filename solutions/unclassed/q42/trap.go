package q42

func trap(height []int) int {
	n :=  len(height)
	if n < 1 {
		return 0
	}
	l, r := 0, n-1
	lMax, rMax := height[l], height[r]
	ans := 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])

		if lMax < rMax {
			ans += lMax - height[l]
			l++
		} else {
			ans += rMax - height[r]
			r--
		}
	}
	return ans
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}