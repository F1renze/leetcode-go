package medium

import "strings"

/**
1. 快慢针需要考虑区间内重复的问题
 */
func lengthOfLongestSubstring3(s string) (maxLength int) {
	n := len(s)
	if n < 1 {
		return
	}
	l := 1
	on := true
	for fast, slow := 1, 0; fast < n && (n-fast-1) > maxLength; {
		if s[fast] != s[slow] {
			on = true
			fast++
			l++
		} else if !on {
			fast++
			slow++
		} else {
			slow = fast
			fast++
			if l > maxLength {
				maxLength = l
			}
			l = 1
		}

	}
	return
}

/**
Contain 只给 tf，无法确定中间那个位置重复
 */
func lengthOfLongestSubstring2(s string) (maxLength int) {
	n := len(s)
	if n < 1 {
		return
	}
	l, r := 0, 1
	sub, t := s[l:r], ""
	maxLength = 1
	for r < n {
		t = string(s[r])
		if strings.Contains(sub, t) || strings.Compare(sub, t) == 0 {
			l++
			if r+1 < n-1 && s[l] == s[r] {
				l++
				r++
			}
		}
		r++
		sub = s[l:r]
		if len(sub) > maxLength {
			maxLength = len(sub)
		}
	}
	return
}

/**
使用 idx 取得重复的位置，将窗口移动到这之后
 */
func lengthOfLongestSubstring(s string) (maxLen int) {
	var l, r, idx, length int
	for i := range s {
		idx = strings.Index(s[l:i], string(s[i]))
		if idx == -1 {
			r++
		} else {
			// idx 相对于子串取得
			l += idx + 1
			r = i + 1
		}
		if length = r-l; length > maxLen {
			maxLen = length
		}
	}
	return
}
