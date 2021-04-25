package q912

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	testCases := []struct {
		give []int
		want []int
	}{
		{
			[]int{5, 1, 1, 2, 0, 0},
			[]int{0, 0, 1, 1, 2, 5},
		},
		{
			[]int{1, 2, 3, 5},
			[]int{1, 2, 3, 5},
		},
		{
			[]int{6, 1, 2, 7, 9, 3, 4, 5, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		got := QuickSortWrap(tc.give)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("want %v, got %v", tc.want, got)
		}
	}

}
