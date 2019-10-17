package q119

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	want := []int{
		1, 4, 6, 4, 1,
	}

	got := GetRow(4)

	if reflect.DeepEqual(want, got) == false {
		t.Fatalf("got '%v'", got)
	}
}
