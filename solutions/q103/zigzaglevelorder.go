package q103

import "github.com/f1renze/leetcode-go/dskit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *dskit.TreeNode) [][]int {
	var (
		result  [][]int
		reverse bool
	)
	if root == nil {
		return result
	}
	queue := []*dskit.TreeNode{root}

	for {
		if len(queue) < 1 {
			break
		}
		tmp := make([]int, 0)
		nextLayer := make([]*dskit.TreeNode, 0)

		for _, node := range queue {
			if !reverse {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}

			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}

		queue = nextLayer
		result = append(result, tmp)
		reverse = !reverse
	}

	return result
}
