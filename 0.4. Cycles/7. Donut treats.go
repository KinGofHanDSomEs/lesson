package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		if i%3 != 0 && i%5 != 0 {
			sum += i
		}
	}
	fmt.Print(sum)
}
