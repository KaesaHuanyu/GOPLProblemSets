package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now() //程序开始时间
	ch1 := make(chan string)
	ch2 := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch1) // 开始一个goroutine
		go fetch(url, ch2)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now() //goroutine 开始时间
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) //只需要io.Copy返回的字节数，不在意resp.Body的内容，可以将其写入ioutil.Discard（可以视为垃圾桶）
	resp.Body.Close()                                 //防止资源泄露
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}
