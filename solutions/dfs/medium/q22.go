package medium


/**
使用闭包 dfs + 剪枝（回溯）
 */
func generateParenthesis(n int) (ans []string) {
	var dfs func(l, r int, path string)
	dfs = func(l, r int, path string) {
		if len(path) == 2*n {
			ans = append(ans, path)
			return
		}
		if l > 0 {
			dfs(l-1, r, path+"(")
		}
		// 回溯
		if l < r {
			dfs(l, r-1, path+")")
		}
	}
	dfs(n, n, "")
	return
}