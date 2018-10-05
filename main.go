// vaportext - create ｖａｐｏｒｗａｖｅ text
package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var version = "to be set by makefile"

func main() {
	input := struct {
		SendToClip bool   `survey:"clip"`
		ZalgoInput bool   `survey:"zalgo"`
		Text       string `survey:"text"`
	}{}

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
