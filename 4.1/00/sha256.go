package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("lI960127."))
	c2 := sha256.Sum256([]byte("lI960127."))
	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}
