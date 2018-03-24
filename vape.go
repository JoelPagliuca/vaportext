package main

import (
	"strings"
)

const fullwidthOffset int32 = 0xFEE0

// vapor - convert to vaportext
func vapor(s string) string {
	var buffer strings.Builder
	for _, c := range s {
		if c == ' ' {
			buffer.WriteRune(c)
		} else {
			buffer.WriteRune(c + fullwidthOffset)
		}
	}
	return buffer.String()
}
