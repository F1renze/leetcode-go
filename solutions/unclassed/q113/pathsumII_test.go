package q113

import (
	"testing"

	"github.com/f1renze/leetcode-go/dskit"
)

func TestStack(t *testing.T) {
	var want []int
	s := NewStack()
	s.push(1)
	s.push(2)
	s.pop()
	want = append(want, 1)
	dskit.Equals(t, want, s.toSlice())

	s.push(3)
	want = append(want, 3)
	dskit.Equals(t, want, s.toSlice())

	r := make([][]int, 0)
	r = append(r, s.toSlice())
	dskit.Equals(t, [][]int{want}, r)
}

func TestPathSum(t *testing.T) {
	tc := []int{
		1, 0, 1, 1, 2, 0, -1, 0, 1, -1, 0, -1, 0, 1, 0,
	}
	root := dskit.BuildTreeFromArr(tc)
	sum := 2
	want := [][]int{
		{1, 0, 1, 0},
		{1, 0, 2, -1},
		{1, 1, 0, 0},
		{1, 1, -1, 1},
	}

	got := PathSum(root, sum)
	dskit.Equals(t, want, got)
}
