package q452

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) < 1 {
		return 0
	}
	sort.Slice(points, func (a, b int) bool {
		// 按尾部排序
		return points[a][1] < points[b][1]
	})

	s := points[0][1]
	c := 1
	for i := range points {
		// 没有共同区间
		if points[i][0] > s {
			c++
			s = points[i][1]
		}
	}
	return c
}