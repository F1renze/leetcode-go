package q881

import "sort"

func numRescueBoats(people []int, limit int) int {
	if len(people) < 1 {
		return 0
	}
	sort.Ints(people)

	s, e := 0, len(people)-1
	c := 0

	for s <= e {
		if people[s]+people[e] <= limit {
			s++
			e--
		} else if people[e] <= limit {
			e--
		}
		c++
	}
	return c
}
