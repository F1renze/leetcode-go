package q121


func maxProfit(prices []int) int {
	n := len(prices) - 1
	if n < 1 {
		return 0
	}
	min, max := 1 << 31, 0

	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i] - min > max {
			max = prices[i] - min
		}
	}

	return max
}