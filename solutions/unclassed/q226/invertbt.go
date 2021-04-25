package q226

import "github.com/f1renze/leetcode-go/dskit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *dskit.TreeNode) *dskit.TreeNode {
	if root == nil {
		return nil
	}
	newLeft := invertTree(root.Right)
	newRight := invertTree(root.Left)

	root.Left = newLeft
	root.Right = newRight
	return root
}
