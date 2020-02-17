package q31

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j > 0 && nums[j] <= nums[i] {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	reverseArr(nums, i+1)
}

func reverseArr(nums []int, lo int) {
	hi := len(nums) - 1
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
}
