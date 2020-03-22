package q572

import "github.com/f1renze/leetcode-go/dskit"

func isSubtree(s *dskit.TreeNode, t *dskit.TreeNode) bool {
	if t == nil && s == nil {
		return true
	}
	if t == nil || s == nil {
		return false
	}

	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(s *dskit.TreeNode, t *dskit.TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if t == nil || s == nil {
		return false
	}
	return s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
