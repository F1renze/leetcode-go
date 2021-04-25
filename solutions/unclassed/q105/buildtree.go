package q105

import "github.com/f1renze/leetcode-go/dskit"

/**
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

func buildTree(preorder []int, inorder []int) *dskit.TreeNode {

	if len(inorder) < 1 {
		return nil
	}

	v := preorder[0]
	root := &dskit.TreeNode{Val: v}

	pos := Index(inorder, v)

	left := inorder[:pos]
	right := inorder[pos+1:]

	root.Left = buildTree(preorder[1:pos+1], left)
	root.Right = buildTree(preorder[pos+1:], right)

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
