package dskit

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func BuildTreeFromArr(arr []int) *TreeNode {
	size := len(arr)
	if size == 0 {
		return nil
	}

	var queue []*TreeNode
	root := NewTree(arr[0])
	queue = append(queue, root)

	for i := 1; i < size; i++ {
		node := queue[0]
		queue = queue[1:]

		node.Left = NewTree(arr[i])
		queue = append(queue, node.Left)
		i++

		if i < size {
			node.Right = NewTree(arr[i])
			queue = append(queue, node.Right)
		}
	}

	return root
}
