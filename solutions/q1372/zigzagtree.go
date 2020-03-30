package q1372

import "github.com/f1renze/leetcode-go/dskit"

// DFS Solution 1
var ans int

func longestZigZag(root *dskit.TreeNode) int {
	ans = 0
	if root == nil {
		return ans
	}

	helper(root, 0, 0)
	helper(root, 1, 0)

	return ans
}

// 0 -> left, 1 -> right
func helper(node *dskit.TreeNode, isLeft, length int) {
	if node == nil {
		return
	}
	ans = max(ans, length)

	if isLeft < 1 {
		helper(node.Left, 0, 1)
		helper(node.Right, 1, length+1)
	} else {
		helper(node.Left, 0, length+1)
		helper(node.Right, 1, 1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS Solution 2

// 0 -> left, 1 -> right
func helper2(node *dskit.TreeNode, isLeft int) int {
	if node == nil {
		return 0
	}
	l, r := helper2(node.Left, 0), helper2(node.Right, 1)
	ans = max(ans, max(l, r))

	if isLeft < 1 {
		return r + 1
	}
	return l + 1
}
