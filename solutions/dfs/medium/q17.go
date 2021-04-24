package medium


import "strconv"

/**
思路同 22 , 多路 dfs
 */

func letterCombinations(digits string) (ans []string) {
	n := len(digits)
	if n < 1 {
		return
	}

	m := genDigitsMap()

	var dfs func(digits, path string, idx int)
	dfs = func(digits, path string, idx int) {
		if len(path) == n {
			ans = append(ans, path)
			return
		}
		sub := m[string(digits[idx])]
		for i := 0; i < len(sub); i++ {
			dfs(digits, path+string(sub[i]), idx+1)
		}
	}

	dfs(digits, "", 0)
	return
}

func genDigitsMap() map[string]string {
	s := genLowercaseAlphabets()
	m := make(map[string]string, 8)
	step := 0
	for i, j := 2, 0; i < 10; i++ {
		if i == 7 || i == 9 {
			step = 4
		} else {
			step = 3
		}
		m[strconv.Itoa(i)] = s[j:j+step]
		j+= step
	}
	return m
}

func genLowercaseAlphabets() (s string) {
	for i := 0; i < 26; i++ {
		s += string('a'+i)
	}
	return
}
