// vaportext - create ｖａｐｏｒｗａｖｅ text
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var version = "to be set by makefile"

func main() {
	interactive := flag.Bool("i", false, "interactive")
	sendToClip := flag.Bool("c", false, "Copy output to clipboard")
	zalgoInput := flag.Bool("z", false, "Make the text Zalgo")
	flag.Parse()

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
