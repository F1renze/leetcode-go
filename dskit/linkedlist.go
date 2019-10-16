package dskit

type ListNode struct {
	Val int
	Next *ListNode
}

func SliceByNode(head *ListNode) (r []int) {
	for {
		if head == nil {
			break
		}
		r = append(r, head.Val)
		head = head.Next
	}
	return
}

func NewLinkedListBySlice(give []int) *ListNode {
	var prev, head *ListNode
	for _, i := range give {
		if head == nil {
			head = &ListNode{Val: i, Next: nil}
			prev = head
			continue
		}
		prev.Next = &ListNode{Val: i, Next: nil}
		prev = prev.Next
	}
	return head
}

func NodeAndSliceEqual(a *ListNode, b []int) bool {
	aVals := SliceByNode(a)

	return Equal(aVals, b)
}

func NodeEqual(a *ListNode, b *ListNode) bool {
	aVals := SliceByNode(a)
	bVals := SliceByNode(b)

	return Equal(aVals, bVals)
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
