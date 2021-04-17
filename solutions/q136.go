package solutions

/**
136. 只出现一次的数字
 */
func singleNumber(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	return sum
}