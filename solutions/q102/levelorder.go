package q102

import "github.com/f1renze/leetcode-go/dskit"

func LevelOrder(root *dskit.TreeNode) [][]int {
	var (
		r [][]int
		q []*dskit.TreeNode
	)

	if root != nil {
		q = append(q, root)
	}

	for {
		if len(q) < 1 {
			break
		}

		var t []int
		var nq []*dskit.TreeNode

		for _, n := range q {
			t = append(t, n.Val)

			if n.Left != nil {
				nq = append(nq, n.Left)
			}

			if n.Right != nil {
				nq = append(nq, n.Right)
			}
		}

		r = append(r, t)
		q = nq
	}
	return r
}
