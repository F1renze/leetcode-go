package q106

import "github.com/f1renze/leetcode-go/dskit"

/**
Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

func buildTree(inorder []int, postorder []int) *dskit.TreeNode {

	if len(inorder) < 1 {
		return nil
	}

	v := postorder[len(postorder)-1]
	root := &dskit.TreeNode{Val: v}

	pos := Index(inorder, v)

	left := inorder[:pos]
	right := inorder[pos+1:]

	root.Left = buildTree(left, postorder[:pos])
	root.Right = buildTree(right, postorder[pos:len(postorder)-1])

	return root
}

func Index(slice []int, item int) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}

	return -1
}
