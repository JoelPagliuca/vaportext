package main

import (
	"math/rand"
	"strings"
	"time"
)

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
			count := rand.Intn(5)
			for i := 0; i < count; i++ {
				buffer.WriteRune(rune(rand.Intn(879-768)) + 768)
			}
		}
	}
	return buffer.String()
}
