package main

import (
	"fmt"
)

func main() {
	fmt.Println(SameString("aassdd", "dasdsa"))
}

func SameString(s, p string) bool {
	if len(s) != len(p) {
		return false
	}
	a := []byte(s)
	b := []byte(p)
	for _, v := range a {
		for x, y := range b {
			if v == y {
				b[x] = b[len(b)-1]
				b = b[:len(b)-1]
				break
			}
		}
	}
	if len(b) == 0 {
		return true
	}
	return false
}
