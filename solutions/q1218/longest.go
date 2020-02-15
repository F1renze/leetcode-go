package q1218

// my solution, inspired by "dp[i] = 1 + dp[i-k]"
func longestSubsequence(arr []int, difference int) int {
	if len(arr) <= 1 {
		return len(arr)
	}
	dict := map[int]int{}
	t, m := 0, 1

	for i := range arr {

		t = arr[i] - difference

		if _, ok := dict[t]; ok {
			dict[arr[i]] = 1 + dict[t]
			m = max(m, dict[arr[i]])
		} else {
			dict[arr[i]] = 1
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}