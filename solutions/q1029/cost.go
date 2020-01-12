package q1029

import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	var c int
	l := len(costs) / 2
	for i := range costs {
		if i < l {
			c += costs[i][0]
		} else {
			c += costs[i][1]
		}
	}
	return c
}