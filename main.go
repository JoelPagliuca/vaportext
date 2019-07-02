// vaportext - create ｖａｐｏｒｗａｖｅ text
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var version = "go-installed"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [...]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "version: %s\n", version)
	}
	interactive := flag.Bool("i", false, "Interactive")
	sendToClip := flag.Bool("c", false, "Copy output to clipboard")
	zalgoInput := flag.Bool("z", false, "Make the text Zalgo")
	versionFlag := flag.Bool("v", false, "Print version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	inputText := flag.Args()
	var buffer strings.Builder
	var space string
	for _, v := range inputText {
		buffer.WriteString(space)
		buffer.WriteString(v)
		space = " "
	}

	input := struct {
		SendToClip bool   `survey:"clip"`
		ZalgoInput bool   `survey:"zalgo"`
		Text       string `survey:"text"`
	}{
		SendToClip: *sendToClip,
		ZalgoInput: *zalgoInput,
		Text:       buffer.String(),
	}

	if *interactive {
		survey.Ask([]*survey.Question{
			{
				Name:   "clip",
				Prompt: &survey.Confirm{Message: "Send to clip?"},
			},
			{
				Name:   "zalgo",
				Prompt: &survey.Confirm{Message: "Zalgo?"},
			},
			{
				Name:   "text",
				Prompt: &survey.Input{Message: "Input:"},
			},
		}, &input)
	}

	output := input.Text
	if input.ZalgoInput {
		output = zalgo(output)
	} else {
		output = vapor(output)
	}

	if input.SendToClip {
		clipboard.WriteAll(output)
	} else {
		fmt.Printf("%s\n", output)
	}
}
