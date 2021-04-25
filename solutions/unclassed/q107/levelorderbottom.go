package q107

import "github.com/f1renze/leetcode-go/dskit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *dskit.TreeNode) (result [][]int) {
	if root == nil {
		return
	}
	queue := []*dskit.TreeNode{root}

	for {
		if len(queue) < 1 {
			break
		}
		tmp := make([]int, 0)
		nextLayer := make([]*dskit.TreeNode, 0)

		for _, node := range queue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		queue = nextLayer
		result = append([][]int{tmp}, result...)
	}

	return
}
