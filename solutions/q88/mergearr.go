package q88


/**
双指针由后往前
 */
func Merge(nums1 []int, m int, nums2 []int, n int)  {

	p1 := m - 1
	p2 := n - 1
	// nums1 尾部
	p := m + n - 1

	for {
		if p1 < 0 || p2 < 0 {
			break
		}

		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for i := 0; i <= p2; i++ {
		nums1[i] = nums2[i]
	}

}