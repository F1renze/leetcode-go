package dskit

import "testing"

func TestEqual(t *testing.T) {
	give := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4}

	if Equal(give, want) == false {
		t.Fatalf("equal failed")
	}
}

/**
测试此项之前必须先测试 NewLinkedListBySlice
*/
func TestSliceByNode(t *testing.T) {
	t.Run("test new linkedlist by slice", testNewLinkedListBySlice)

	give := []int{1, 2, 3, 4}
	head := NewLinkedListBySlice(give)
	got := SliceByNode(head)

	if Equal(give, got) == false {
		t.Fatalf("get slice by node failed, got '%v'", got)
	}

}

func testNewLinkedListBySlice(t *testing.T) {
	give := []int{1, 2, 3, 4}
	gotHead := NewLinkedListBySlice(give)
	var got []int

	cur := gotHead
	for {
		if cur == nil {
			break
		}
		got = append(got, cur.Val)
		cur = cur.Next
	}
	if Equal(got, give) == false {
		t.Fatalf("new list by slice failed, got '%v'", got)
	}
}
