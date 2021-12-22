package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	if N < 10 {
		fmt.Printf("AGC00%d", N)
	} else if N < 42 {
		fmt.Printf("AGC0%d", N)
	} else {
		fmt.Printf("AGC0%d", N+1)
	}
}
