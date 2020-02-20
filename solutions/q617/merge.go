package q617

import "github.com/f1renze/leetcode-go/dskit"

func mergeTrees(t1 *dskit.TreeNode, t2 *dskit.TreeNode) *dskit.TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left,t2.Left)
	t1.Right =mergeTrees(t1.Right,t2.Right)
	return t1
}