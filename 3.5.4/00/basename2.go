package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("a/b/c/d.a.go")
	fmt.Println(basename2("a/b/c/d.a.go"))
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
