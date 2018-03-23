package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, v := range args {
		fmt.Println(v)
	}
}
