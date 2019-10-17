package porb118


func Calc(i, j int, arr [][]int) int {
	if (j == 1) || (j == i) {
		return 1
	}
	qj := j - 1
	qi := i - 1
	if qi >= 0 && qi < len(arr) &&  qj >= 0  && qj < len(arr[qi]) {
		return arr[qi][qj]
	}

	return Calc(i-1, j-1, arr) + Calc(i-1, j, arr)
}

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
