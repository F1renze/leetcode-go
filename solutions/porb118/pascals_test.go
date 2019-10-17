package porb118

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	want := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

	got := Generate(5)

	if reflect.DeepEqual(want, got) == false {
		t.Fatalf("got '%v'", got)
	}
}

