package q299

import "strconv"

func getHint(secret string, guess string) string {
	a, b := 0, 0
	mS, mG := make([]int, 10), make([]int, 10)

	for i := range secret {
		cS, cG := secret[i], guess[i]
		if cS == cG {
			a++
		} else {
			mS[cS-'0']++
			mG[cG-'0']++
		}
	}
	for i := range mS {
		b += min(mS[i], mG[i])
	}

	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}