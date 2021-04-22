package medium


func swap(a, b *int) {
	*a, *b = *b, *a
}


/**
48. 旋转图像
给定一个 n×n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
https://leetcode-cn.com/problems/rotate-image
 */


func rotate(matrix [][]int)  {
	row := len(matrix)
	var i, j int
	// 水平翻转
	for i = 0; i < row/2; i++ {
		for j = 0; j < row; j++ {
			swap(&matrix[i][j], &matrix[row-i-1][j])
		}
	}
	// 对角线翻转
	for i = 0; i < row-1; i++ {
		for j = i+1; j < row; j++ {
			swap(&matrix[i][j], &matrix[j][i])
		}
	}
}