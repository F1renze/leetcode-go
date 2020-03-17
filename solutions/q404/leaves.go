package q404

import "github.com/f1renze/leetcode-go/dskit"

func sumOfLeftLeaves(root *dskit.TreeNode) int {
	ret := 0
	if root == nil {
		return ret
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ret += root.Left.Val
	}

	ret += sumOfLeftLeaves(root.Left)
	ret += sumOfLeftLeaves(root.Right)

	return ret
}