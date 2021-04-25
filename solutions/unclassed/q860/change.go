package q860

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for i := range bills {
		switch bills[i] {
		case 5:
			five++
		case 10:
			ten++
			five--
		default:
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}

		if five < 0 {
			return false
		}
	}
	return true
}
