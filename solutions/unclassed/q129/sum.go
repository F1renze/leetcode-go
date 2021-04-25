package q129

import "github.com/f1renze/leetcode-go/dskit"

func sumNumbers(root *dskit.TreeNode) (ret int) {
	dfs(root, 0, &ret)
	return
}

func dfs(node *dskit.TreeNode, num int, ret *int) int {
	if node == nil {
		return num
	}
	num = num * 10 + node.Val
	if node.Left == nil && node.Right == nil {
		*ret += num
		return num
	}
	dfs(node.Left, num, ret)
	dfs(node.Right, num, ret)
	return num
}
