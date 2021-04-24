package medium

import (
	"fmt"
	"strings"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	ans := LetterCombinations("23")
	expect := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	if len(ans) != len(expect) {
		t.Fatal("length not match")
	}

	if !sameStringSlice(ans, expect) {
		fmt.Println(strings.Join(expect, ","))
		fmt.Println(strings.Join(ans, ","))
		t.Fatal("ans not match")
	}

}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
