package q55

import "github.com/f1renze/leetcode-go/dskit"

func maxDepth(root *dskit.TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := maxDepth(root.Left), maxDepth(root.Right)

	return 1 + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}