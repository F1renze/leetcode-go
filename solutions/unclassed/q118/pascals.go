package q118

/**
q118. Pascal's Triangle

Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func Calc(i, j int, arr [][]int) int {
	if (j == 1) || (j == i) {
		return 1
	}
	qj := j - 1
	qi := i - 1
	if qi >= 0 && qi < len(arr) && qj >= 0 && qj < len(arr[qi]) {
		return arr[qi][qj]
	}

	return Calc(i-1, j-1, arr) + Calc(i-1, j, arr)
}

/**
递归缓存解法
*/
func Generate(numRows int) [][]int {
	var r [][]int

	for i := 1; i <= numRows; i++ {
		var tmp []int
		for j := 1; j <= i; j++ {
			v := Calc(i, j, r)
			tmp = append(tmp, v)
		}
		r = append(r, tmp)
	}
	return r
}
