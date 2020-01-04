package q1005

import "sort"

// 4ms
// todo min heap
func largestSumAfterKNegations(A []int, K int) (r int) {
	sort.Ints(A)
	i, p := 0, 0
	for K > 0 {
		p = A[i]
		A[i] = -A[i]
		// 若取反后值未变下标不变
		// A[i] != p 加速
		// 若比后一个值小下标依然不变
		if A[i] != p && i < len(A)-1 && A[i] >= A[i+1]{
			i++
		}
		K--
	}

	for i := range A {
		r += A[i]
	}
	return r
}
