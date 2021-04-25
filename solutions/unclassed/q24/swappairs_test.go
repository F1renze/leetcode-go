package q24

import (
	"github.com/f1renze/leetcode-go/dskit"
	"testing"
)

func TestSwapPairs(t *testing.T) {

	give := []int{1, 2, 3, 4}
	want := []int{2, 1, 4, 3}

	head := dskit.NewLinkedListBySlice(give)

	newHead := SwapPairs(head)

	if eq := dskit.NodeAndSliceEqual(newHead, want); eq == false {
		t.Fatalf("test failed, want '%v' != '%v",
			want, dskit.SliceByNode(newHead))
	}
}
