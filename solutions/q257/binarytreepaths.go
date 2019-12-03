package q257

import (
	"github.com/f1renze/leetcode-go/dskit"
	"strconv"
)

/**
int search(int t)
{
    if(满足输出条件)
    {
        输出解;
    }
    else
    {
        for(int i=1;i<=尝试方法数;i++)
            if(满足进一步搜索条件)
            {
                为进一步搜索所需要的状态打上标记;
                search(t+1);
                恢复到打标记前的状态;//也就是说的{回溯一步}
            }
    }
}
*/

func DepthFirstSearch(root *dskit.TreeNode) []string {
	var (
		result []string
		helper func(*dskit.TreeNode, string, *[]string)
	)

	helper = func(node *dskit.TreeNode, val string, paths *[]string) {
		if node == nil {
			return
		}
		val += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			*paths = append(*paths, val)
		} else {
			val += "->"
			helper(node.Left, val, paths)
			helper(node.Right, val, paths)
		}
	}

	helper(root, "", &result)

	return result
}
