package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status) //resp.Status是个结构体，但是定义了String方法，
	} //实现了fmt.stringer接口，所以可以直接通过fmt.Println等函数打印输出
}
