package q820

import (
	"sort"
	"strings"
)

/**
更高效的方法, 在 map 中查询子字符串
 */
func minimumLengthEncodingFaster(words []string) int {
	m := make(map[string]interface{}, len(words))

	for i := range words {
		m[words[i]] = struct{}{}
	}

	for _, w := range words {
		for i := 1; i < len(w);i++ {
			v := w[i:]
			if _, ok := m[v]; ok {
				delete(m, v)
			}
		}
	}
	c := 0
	for k, _ := range m {
		c += len(k) + 1
	}
	return c
}

/**
Trie 思路, 逆序后按字典序排序, 前一单词为后一单词前缀的跳过
em
emit
lleb
 */
func minimumLengthEncoding(words []string) int {
	n := len(words)
	rws := make([]string, 0, n)
	for _, w := range words {
		rws = append(rws, Reverse(w))
	}

	sort.Slice(rws, func(a, b int)bool{
		return rws[a] < rws[b]
	})
	c := 0

	for i, w := range rws {
		if i + 1 < n && strings.HasPrefix(rws[i+1], w) {
			continue
		}
		c += len(w) + 1
	}
	return c
}

func Reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}