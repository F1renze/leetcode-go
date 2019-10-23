package q94

import "../../dskit"

func inorderTraversal(root *dskit.TreeNode) []int {
	var (
		s []*dskit.TreeNode
		r []int
	)

	if root == nil {
		return nil
	}

	for {
		if len(s) < 1 && root == nil {
			break
		}

		if root != nil {
			s = append(s, root)
			root = root.Left
		} else {
			root = s[len(s)-1]
			s = s[:len(s)-1]

			r = append(r, root.Val)
			root = root.Right
		}

	}

	return r
}
