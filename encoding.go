package main

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

const (
	allLower string = "abcdefghijklmnopqrstuvwxyz"
	allUpper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	fullwidthOffset        int32 = 0xFEE0
	frakturLowerCaseOffset int32 = 0x1D4BD
	frakturUpperCaseOffset int32 = 0x1D52B
)

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

// Range for Combining characters
// 768,879
func zalgo(s string) string {
	rand.NewSource(time.Now().UnixNano())
	var buffer strings.Builder
	for _, c := range s {
		if c == ' ' {
			buffer.WriteRune(c)
		} else {
			buffer.WriteRune(c)
			count := rand.Intn(10)
			for i := 0; i < count; i++ {
				buffer.WriteRune(rune(rand.Intn(879-768)) + 768)
			}
		}
	}
	return buffer.String()
}

func fraktur(s string) string {
	var buffer strings.Builder
	for _, c := range s {
		if unicode.IsLower(c) {
			buffer.WriteRune(c + frakturLowerCaseOffset)
		} else if unicode.IsUpper(c) {
			buffer.WriteRune(c + frakturUpperCaseOffset)
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}
