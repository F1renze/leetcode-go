package shuffle

import (
	"fmt"
	"testing"
)

func TestFisherYatesShuffle(t *testing.T) {
	tc := [][]int {
		{10, 9, 8, 7, 6, 5, 4, 3},
	}

	for i := range tc {
		fmt.Println(FisherYatesShuffle(tc[i]))
	}
}
