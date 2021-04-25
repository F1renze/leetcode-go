package q04_02

import "github.com/f1renze/leetcode-go/dskit"

func sortedArrayToBST(nums []int) *dskit.TreeNode {
	return helper(nums)
}

func helper(nums []int) *dskit.TreeNode {
	if len(nums) < 1 {
		return nil
	}

	mid := len(nums) >> 1
	node := &dskit.TreeNode{
		Val: nums[mid],
	}

	node.Left = helper(nums[:mid])
	node.Right = helper(nums[mid+1:])

	return node
}