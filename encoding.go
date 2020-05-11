package main

import (
	"math/rand"
	"strings"
	"time"
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
