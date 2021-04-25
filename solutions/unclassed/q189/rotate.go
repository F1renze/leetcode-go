package q189

// 暴力
func rotateBruteForce(nums []int, k int) {
	l := len(nums)
	if l < 2 {
		return
	}

	var t int
	for i := 0; i < k; i++ {
		t = nums[l-1]
		for j := range nums {
			nums[j], t = t, nums[j]
		}
	}
}

// 反转
func rotateReverse(nums []int, k int) {
	l := len(nums)
	if l < 2 {
		return
	}

	reverseArr(nums, 0, l-1)
	reverseArr(nums, 0, (k%l)-1)
	reverseArr(nums, (k % l), l-1)
}

func reverseArr(nums []int, lo, hi int) {
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]

		lo++
		hi--
	}
}

// 额外空间
func rotateOn(nums []int, k int) {
	l := len(nums)
	if l < 2 {
		return
	}

	r := make([]int, l)
	for i := range nums {
		r[(i+k)%l] = nums[i]
	}

	for i := range nums {
		nums[i] = r[i]
	}
}

// todo 环状替换
