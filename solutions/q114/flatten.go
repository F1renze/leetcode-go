package q114

import "github.com/f1renze/leetcode-go/dskit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *dskit.TreeNode) {
	if root == nil {
		return
	}
	var (
		s   []*dskit.TreeNode
		cur *dskit.TreeNode
		tmp *dskit.TreeNode
	)

	for len(s) > 0 || root != nil {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}

		if len(s) > 0 {
			cur = s[len(s)-1]
			s = s[:len(s)-1]

			tmp = cur.Right
			cur.Right = cur.Left
			cur.Left = nil

			for cur.Right != nil {
				cur = cur.Right
			}

			cur.Right = tmp
			root = tmp
		}
	}
}

// recursion
func recursionFlatten(root *dskit.TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	r := root.Right
	root.Right, root.Left = root.Left, nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = r
}
