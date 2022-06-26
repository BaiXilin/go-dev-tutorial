package main

import (
	"fmt"

	"example/baixilin/hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(morestrings.ReverseRunes("Hello World!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
