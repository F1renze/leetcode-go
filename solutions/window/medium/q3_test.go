package medium

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tc := []struct {
		str    string
		output int
	}{
		{"abba", 2},
		{"au", 2},
		{" ", 1},
		{"pwwkew", 3},
		{"aab", 2},
		{"aa", 1},
		{"bbabcabcbb", 3},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"", 0},

	}
	var length int
	for i := 0; i < len(tc); i++ {
		length = lengthOfLongestSubstring(tc[i].str)
		if length != tc[i].output {
			t.Errorf("\"%s\" expect %d, got %d", tc[i].str, tc[i].output, length)
		}
	}

}
