// vaportext - create ｖａｐｏｒｗａｖｅ text
package main

import (
	"fmt"
	"os"
	"strings"
)

var version = "to be set by makefile"

func main() {
	// read in the string from input
	var buffer strings.Builder
	args := os.Args[1:]
	// combine into single string
	var space string
	for _, v := range args {
		buffer.WriteString(space)
		buffer.WriteString(v)
		space = " "
	}
	s := vapor(buffer.String())
	fmt.Printf("%s\n", s)
}
