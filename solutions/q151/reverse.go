package q151

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	g := strings.Split(s, " ")
	t := make([]string, 0, len(g))
	for i := len(g) - 1; i >= 0; i-- {
		if len(strings.TrimSpace(g[i])) < 1 {
			continue
		}
		t = append(t, g[i])
	}
	return strings.Join(t, " ")
}

