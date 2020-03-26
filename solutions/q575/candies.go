package q575


func distributeCandies(candies []int) int {
	n := len(candies) / 2
	s := make(map[int]struct{})
	for _, i := range candies {
		if _, ok := s[i]; !ok {
			s[i] = struct{}{}
		}
	}
	categories := len(s)
	if categories >= n {
		return n
	}
	return categories
}

