package q119

import "fmt"

/**
119. Pascal's Triangle II

Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 3
Output: [1,3,3,1]
Follow up:

Could you optimize your algorithm to use only O(k) extra space?
*/

/**
递归，空间 O(k) 解法
*/
func GetRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prevRow := GetRow(rowIndex - 1)
	var row []int

	for i := 0; i < rowIndex+1; i++ {
		if i == 0 || i == rowIndex {
			row = append(row, 1)
		} else {
			row = append(row, prevRow[i-1]+prevRow[i])
		}
	}
	fmt.Printf("%v row %v\n", rowIndex, row)
	return row
}
