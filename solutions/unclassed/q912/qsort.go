package q912

func QuickSortWrap(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	QuickSort(arr, 0, len(arr)-1)
	return arr
}

func QuickSort(arr []int, j, k int) {
	if j >= k {
		return
	}

	// 基准值
	t := arr[j]
	l, r := j, k

	for {
		if l == r {
			break
		}

		for {
			if l >= r || arr[r] < t {
				break
			}
			r--
		}

		for {
			if l >= r || arr[l] > t {
				break
			}
			l++
		}

		if l < r {
			Swap(arr, l, r)
		}
	}

	arr[j] = arr[l]
	arr[l] = t

	QuickSort(arr, j, l-1)
	QuickSort(arr, l+1, k)
}

func Swap(arr []int, j, k int) {
	tmp := arr[j]
	arr[j] = arr[k]
	arr[k] = tmp
}
