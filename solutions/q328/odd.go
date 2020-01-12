package q328

import "github.com/f1renze/leetcode-go/dskit"

func oddEvenList(head *dskit.ListNode) *dskit.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next

	// 偶链表头
	t := even

	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = t

	return head
}
