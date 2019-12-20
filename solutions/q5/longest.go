package q5

import "strings"

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return s
	}
	sArr := strings.Split(s, "")
	s = "#" + strings.Join(sArr, "#") + "#"

	l := make([]int, len(s))
	for i := range l {
		l[i] = 1
	}

	var r, c, m int

	for i := 1; i < len(l); i++ {

		if len(l) - c <= l[m]{
			break
		}

		// 在边界内时，当镜像i 回文半径不超过边界直接 copy
		if i < r {
			l[i] = min(r-i, l[2*c-i])
		}

		for i-l[i] >= 0 && i+l[i] < len(l) && s[i-l[i]] == s[i+l[i]] {
			// 扩散计数
			l[i]++
		}

		if i + l[i] > r {
			r = i+l[i]
			c = i
		}
		if l[i] > l[m] {
			m = i
		}
	}
	// l[i]++ 处跳出循环前多加了 1
	l[m]--

	return strings.Replace(s[m-l[m]:m+l[m]], "#", "", -1)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}