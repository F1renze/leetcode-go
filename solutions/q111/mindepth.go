package q111

import "github.com/f1renze/leetcode-go/dskit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *dskit.TreeNode) (depth int) {
	if root == nil {
		return
	}

	queue := []*dskit.TreeNode{root}

	for {
		if len(queue) < 1 {
			break
		}
		depth++
		nextLayer := make([]*dskit.TreeNode, 0)

		for _, node := range queue {
			if node.Left == node.Right && node.Left == nil {
				return
			}
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		queue = nextLayer
	}
	return
}
