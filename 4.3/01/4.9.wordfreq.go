package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	file := os.Args[1]
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)

	input.Split(bufio.ScanWords) //不需要参数

	for input.Scan() {
		words[input.Text()]++
	}
	fmt.Println(words)
}
