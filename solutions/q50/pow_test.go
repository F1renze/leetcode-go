package q50

import "testing"

func TestFastPow(t *testing.T) {
	want := float64(1024)
	got := Pow(2.0, 10)

	if got != want {
		t.Fatalf("fast pow failed, '%v'", got)
	}
}
