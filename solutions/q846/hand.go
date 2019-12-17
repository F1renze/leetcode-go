package q846


func isNStraightHand(hand []int, W int) bool {
	candidate := make([]int, W)
	for i := range hand {
		candidate[hand[i] % W]++
	}

	for i := 0; i < W-1; i++ {
		// 若能被分为 w 个顺子，每个索引的值相同
		if candidate[i] != candidate[i+1] {
			return false
		}
	}
	return true
}

