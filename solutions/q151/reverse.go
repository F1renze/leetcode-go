package q151

import "strings"

func reverseWords(s string) string {
	g := strings.Split(s, " ")
	r := []string{}
	for i := len(g)-1; i >= 0; i-- {
		if g[i] == "" {
			continue
		}
		r = append(r, g[i])
	}
	return strings.Join(r, " ")
}

