package q145

import "github.com/f1renze/leetcode-go/dskit"

func postorderTraversal(root *dskit.TreeNode) []int {

	var (
		r []int
		s []*dskit.TreeNode
	)

	for {
		if root == nil && len(s) < 1 {
			break
		}

		if root != nil {
			s = append(s, root)
			// insert val at index zero
			r = append([]int{root.Val}, r...)

			root = root.Right
		} else {
			root = s[len(s)-1].Left
			s = s[:len(s)-1]

		}
	}

	return r

}
