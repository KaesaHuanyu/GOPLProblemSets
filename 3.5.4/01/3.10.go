package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, j := n%3, 0; i <= n; i, j = i+3, i {
		buf.WriteString(s[j:i] + ",")
	}
	return buf.String()[:len(buf.String())-1]
}
