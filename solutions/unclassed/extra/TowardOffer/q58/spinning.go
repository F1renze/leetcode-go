package q58


func reverseLeftWords(s string, n int) string {
	if n >= len(s) {
		return s
	}
	return s[n:] + s[:n]
}