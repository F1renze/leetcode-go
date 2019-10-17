package q206

import "../../dskit"

/**
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

func reverseListIteration(head *dskit.ListNode) *dskit.ListNode {
	var prev *dskit.ListNode
	cur := head

	for {
		if cur == nil {
			break
		}

		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

func reverseListRecursion(head *dskit.ListNode) *dskit.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}
