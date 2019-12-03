package dskit

import (
	"reflect"
	"testing"
)

func Range(s, e int) (r []int) {
	for i := s; i < e; i++ {
		r = append(r, i)
	}
	return
}

func TAssert(tb testing.TB, condition bool, format string, v ...interface{}) {
	tb.Helper()
	if condition != true {
		tb.Fatalf(format, v...)
	}
}

func TNoErr(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("unexpected error occur, '%v'", err)
	}
}

func Equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("excepted '%v', got '%v'", exp, act)
	}
}

