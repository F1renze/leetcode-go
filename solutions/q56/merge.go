package q56


import "sort"

// sort
func merge(intervals [][]int) [][]int {
	if len(intervals) < 1{
		return nil
	}
	sort.Slice(intervals, func(i, j int)bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{
		intervals[0],
	}

	for i := 1; i < len(intervals);i++ {
		end := len(result)-1
		if intervals[i][0] <= result[end][1] {
			a, b := intervals[i][1], result[end][1]
			if a > b {
				result[end][1] = a
			} else {
				result[end][1] = b
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}