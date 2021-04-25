package q1162

type pos struct {
	x, y int
}

func maxDistance(grid [][]int) int {
	l := len(grid)
	q := make([]*pos, 0, l)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				// enqueue
				q = append(q, &pos{i, j})
			}
		}
	}

	w := len(grid[0])
	// no island || full island
	if len(q) == 0 || len(q) == l * w {
		return -1
	}
	d := []int{0, 0, 1, -1}
	c := 0

	for len(q) > 0 {
		tmpQ := make([]*pos, 0, len(q))

		for _, p := range q {
			for i := range d {
				x, y := p.x+d[i], p.y+d[(i+2)%len(d)]
				if x < 0 || y < 0 || x >= l || y >= w || grid[x][y] > 0 {
					continue
				}
				grid[x][y] = 1
				tmpQ = append(tmpQ, &pos{x, y})
			}
		}
		// distance = level cnt
		c++
		q = tmpQ
	}

	return c - 1
}