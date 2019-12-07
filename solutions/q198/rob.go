package q198


func rob(nums []int) int {
	var prev, cur int

	for i := range nums {
		if (nums[i] + prev) > cur {
			prev, cur = cur, prev + nums[i]
		} else {
			prev = cur
		}
	}
	return cur
}