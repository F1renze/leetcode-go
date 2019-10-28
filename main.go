package main

import (
	"fmt"

	"github.com/f1renze/leetcode-go/build"
)

func main() {
	err := build.Run()
	fmt.Printf("%+v\n", err)
}
