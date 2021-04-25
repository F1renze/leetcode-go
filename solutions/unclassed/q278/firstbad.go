package q278

func isBadVersion(version int) bool {
	// 根据不同测试用例有不同实现
	return true
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for {
		if left > right {
			break
		}

		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
