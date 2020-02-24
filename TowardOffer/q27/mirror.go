package q27

import "github.com/f1renze/leetcode-go/dskit"


func mirrorTree(root *dskit.TreeNode) *dskit.TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)

	return root
}