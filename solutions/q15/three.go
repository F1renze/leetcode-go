package q15

import "sort"

func threeSum(nums []int) (ret [][]int) {
	length := len(nums)
	if length < 3 {
		return
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	var (
		l, r, sum int
	)
	for i := 0; i < length; i++ {
		// 由于已排序，若当前元素大于 0 其后必大于 0
		if nums[i] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l = i + 1
		r = length - 1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				// 去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 去重
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return
}
