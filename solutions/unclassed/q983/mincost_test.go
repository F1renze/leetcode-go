package q983

import "testing"

func TestMincostTickets(t *testing.T) {
	daysTc := [][]int{
		{1, 4, 6, 7, 8, 20},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
	}
	costsTc := [][]int{
		{2, 7, 15},
		{2, 7, 15},
	}
	wants := []int{11, 17}

	for i := range daysTc {
		got := MinCostTickets(daysTc[i], costsTc[i])
		if got != wants[i] {
			t.Fatal("algorithm not working")
		}
	}
}
