package q654

import "github.com/f1renze/leetcode-go/dskit"

func constructMaximumBinaryTree(nums []int) *dskit.TreeNode {
	n := len(nums)
	if n < 1 {
		return nil
	} else if n == 1 {
		return &dskit.TreeNode{
			Val: nums[0],
		}
	}

	maxPos := 0
	for i := range nums {
		if nums[maxPos] < nums[i] {
			maxPos = i
		}
	}

	node := &dskit.TreeNode{
		Val: nums[maxPos],
	}
	node.Left = constructMaximumBinaryTree(nums[:maxPos])
	node.Right = constructMaximumBinaryTree(nums[maxPos+1:])
	return node
}
