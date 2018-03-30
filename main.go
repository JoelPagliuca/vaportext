// vaportext - create ｖａｐｏｒｗａｖｅ text
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

var version = "to be set by makefile"

func main() {
	// get input flags
	sendToClip := flag.Bool("c", false, "Copy output to clipboard")
	zalgoInput := flag.Bool("z", false, "Make the text Zalgo")
	flag.Parse()
	// read in the string from input
	inputText := flag.Args()

	// combine into single string
	var buffer strings.Builder
	var space string
	for _, v := range inputText {
		buffer.WriteString(space)
		buffer.WriteString(v)
		space = " "
	}
	output := buffer.String()

	// do the conversion
	if *zalgoInput {
		output = zalgo(output)
	} else {
		output = vapor(output)
	}
	if *sendToClip {
		clipboard.WriteAll(output)
	} else {
		fmt.Printf("%s\n", output)
	}

}
