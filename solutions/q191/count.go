package q191


func hammingWeightFast(num uint32) int {
	var r int
	for num > 0 {
		r++
		num &= num-1
	}
	return r
}

func hammingWeight(num uint32) int {
	var r int
	for num > 0 {
		if (num & 1) > 0 {
			r++
		}
		num >>= 1
	}
	return r
}