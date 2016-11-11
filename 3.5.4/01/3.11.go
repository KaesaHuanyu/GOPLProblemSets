//支持浮点数处理和一个可选正负号的处理
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("+1234556.23333"))
}

func comma(s string) string {
	var h string
	var f string
	if s[0] == '+' || s[0] == '-' {
		h = s[:1]
		s = s[1:]
	}
	for i, v := range s {
		if v == '.' {
			f = s[i:]
			s = s[:i]
		}
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, j := n%3, 0; i <= n; i, j = i+3, i {
		buf.WriteString(s[j:i] + ",")
	}
	return h + buf.String()[:len(buf.String())-1] + f
}
