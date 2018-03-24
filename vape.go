package main

import (
	"strings"
)

// vapor - convert to vaportext
func vapor(s string) string {
	var buffer strings.Builder
	for _, c := range s {
		buffer.WriteRune(c + (0xFF00 - 32))
	}
	return buffer.String()
}
