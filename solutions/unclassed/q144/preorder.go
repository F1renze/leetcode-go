package q144

import "github.com/f1renze/leetcode-go/dskit"

func preorderTraversal(root *dskit.TreeNode) []int {
	var (
		r []int
		s []*dskit.TreeNode
	)
	s = append(s, root)

	for {
		if len(s) < 1 {
			break
		}
		cur := s[len(s)-1]
		s = s[:len(s)-1]

		if cur == nil {
			continue
		}

		r = append(r, cur.Val)

		s = append(s, cur.Right)
		s = append(s, cur.Left)
	}

	return r
}
