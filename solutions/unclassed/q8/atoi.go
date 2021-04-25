package q8

import "strings"

const (
	max = (1 << 31) -1
	min = -1 << 31
)

func myAtoi(str string) (n int) {
	str = strings.TrimSpace(str)
	if len(str) < 1 {
		return
	}
	sign := 1
	// compare rune & char(in string) way 1
	// str[0]: type byte
	switch str[0] {
	case '+':
		str = str[1:]
	case '-':
		sign, str = -1, str[1:]
	default:
	}
	// compare rune & char(in string) way 2
	runes := []rune(str)
	if len(runes) < 1 {
		return
	}
	if runes[0] - '0' > 9 || runes[0] - '0' < 0 {
		return
	}
	tmp := 0
	for _, char := range runes {
		tmp = int(char - '0')
		if tmp > 9 || tmp < 0 {
			break
		}
		n = n * 10 + tmp
		if n > max && sign > 0 {
			return max
		} else if n > max {
			return min
		}
	}

	return sign * n
}

// todo dfa solution