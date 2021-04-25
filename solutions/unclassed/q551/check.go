package q551

import "strings"

func checkRecord(s string) bool {
	a := 0

	for i := 0; i < len(s) && a < 2; i++ {
		if s[i] == 'A' {
			a++
		}
	}
	return a < 2 && strings.Index(s, "LLL") == -1
}