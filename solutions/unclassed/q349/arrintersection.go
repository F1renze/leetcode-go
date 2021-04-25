package q349

func intersection(nums1 []int, nums2 []int) (arr []int) {
	m := map[int]bool{}

	for _, v := range nums1 {
		m[v] = true
	}

	for _, v := range nums2 {
		if m[v] {
			arr = append(arr, v)
			m[v] = false
		}
	}
	return
}
