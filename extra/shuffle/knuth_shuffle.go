package shuffle

import (
	"math/rand"
	"time"
)

func FisherYatesShuffle(arr []int) []int {
	rand.Seed(time.Now().Unix())
	for i := len(arr) - 1; i >= 0; i-- {
		n := rand.Intn(i + 1)
		arr[i], arr[n] = arr[n], arr[i]
	}
	return arr
}
