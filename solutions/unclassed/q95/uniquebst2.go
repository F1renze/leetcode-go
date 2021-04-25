package q95

import "github.com/f1renze/leetcode-go/dskit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func GenerateTrees(n int) []*dskit.TreeNode {
	if n == 0 {
		return nil
	}
	return GenTreesHelper(1, n)
}

func GenTreesHelper(start, end int) []*dskit.TreeNode {
	var trees []*dskit.TreeNode
	if start > end {
		return append(trees, nil)
	}

	for i := start; i <= end; i++ {
		left := GenTreesHelper(start, i-1)
		right := GenTreesHelper(i+1, end)

		for l := range left {
			for r := range right {
				trees = append(trees,
					&dskit.TreeNode{
						Val:   i,
						Left:  left[l],
						Right: right[r],
					})
			}
		}
	}

	return trees
}
