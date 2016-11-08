package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := fizzBuzz(15)
	fmt.Println(num)
}

func fizzBuzz(n int) []string {
	num := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1) >= 15 && (i+1)%15 == 0 {
			num[i] = "FizzBuzz"
		} else if (i+1) >= 3 && (i+1)%3 == 0 && (i+1)%5 != 0 {
			num[i] = "Fizz"
		} else if (i+1) >= 5 && (i+1)%5 == 0 && (i+1)%3 != 0 {
			num[i] = "Buzz"
		} else {
			num[i] = strconv.Itoa(i + 1)
		}
	}
	return num
}
