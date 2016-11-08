package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s += input.Text()
	}
	fmt.Println(basename1(s))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
