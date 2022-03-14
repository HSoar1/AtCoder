package main

import (
	"fmt"
)

func main() {
	var N, A, B, P, Q, R, S int
	fmt.Scanf("%d %d %d", &N, &A, &B)
	fmt.Scanf("%d %d %d %d", &P, &Q, &R, &S)

	for i := P; i <= Q; i++ {
		for j := R; j <= S; j++ {
			if (i - j) == (A - B) {
				fmt.Print("#")
			} else if (i + j) == (A + B) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
