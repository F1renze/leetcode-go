package q21

import (
	"github.com/f1renze/leetcode-go/dskit"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoListsRecursion(l1 *dskit.ListNode, l2 *dskit.ListNode) *dskit.ListNode {
	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	case l1.Val < l2.Val:
		l1.Next = MergeTwoListsRecursion(l1.Next, l2)
		return l1
	default:
		l2.Next = MergeTwoListsRecursion(l1, l2.Next)
		return l2
	}
}

/**
归并排序
*/
func MergeTwoListsIteration(l1 *dskit.ListNode, l2 *dskit.ListNode) *dskit.ListNode {
	head := new(dskit.ListNode)
	prev := head

	for {
		if l1 == nil || l2 == nil {
			break
		}

		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return head.Next
}
