// vaportext - create ｖａｐｏｒｗａｖｅ text
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

var version = "to be set by makefile"

func convertText(toConvert string, doClipboard bool, doZalgo bool) {
	output := toConvert
	// do the conversion
	if doZalgo {
		output = zalgo(output)
	} else {
		output = vapor(output)
	}
	if doClipboard {
		clipboard.WriteAll(output)
	} else {
		fmt.Printf("%s\n", output)
	}
}

func main() {
	// get input flags
	sendToClip := flag.Bool("c", false, "Copy output to clipboard")
	zalgoInput := flag.Bool("z", false, "Make the text Zalgo")
	flag.Parse()
	// read in the string from input
	stdinInfo, _ := os.Stdin.Stat()
	var inputReader *bufio.Reader
	if stdinInfo.Size() > 0 {
		fmt.Println("stdin")
		// read from pipe
		inputReader = bufio.NewReader(os.Stdin)
		for {
			inputText, err := inputReader.ReadString('\n')
			if err != nil && err == io.EOF {
				break
			}
			convertText(inputText, *sendToClip, *zalgoInput)
		}
	} else {
		inputText := strings.Join(flag.Args(), " ")
		convertText(inputText, *sendToClip, *zalgoInput)
	}

}
