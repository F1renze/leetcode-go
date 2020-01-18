package q983

// 状态转移方程
// dp(i) = min(dp(j1) + costs[0], dp(j7) + costs[1], dp(j30) + costs[2])
func MinCostTickets(days []int, costs []int) int {
	d := len(days)
	if d < 1 {
		return 0
	}

	var (
		dp func(int) int
		durations = []int{1, 7, 30}
	)
	memo := make([]int, d)

	dp = func(i int) int {
		if i >= d {
			return 0
		}
		if memo[i] != 0 {
			return memo[i]
		}

		j := i
		ans := 1 << 32
		for k := range durations {
			for j < d && days[j] < days[i] + durations[k] {
				j++
			}

			ans = min(ans, dp(j) + costs[k])
		}

		memo[i] = ans
		return ans
	}

	return dp(0)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
