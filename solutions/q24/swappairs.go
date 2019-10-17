package q24

import (
	"../../dskit"
)

/**
24. Swap Nodes in Pairs
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
递归
*/
func SwapPairs(head *dskit.ListNode) *dskit.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 取指针值 T: ListNode
	tmp := *head.Next

	tmp.Next = head

	head.Next = SwapPairs(head.Next.Next)

	// 取值地址 T: *ListNode
	return &tmp
}
