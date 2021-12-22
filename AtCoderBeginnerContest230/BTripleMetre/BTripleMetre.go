package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	T1ok := true
	T2ok := true
	T3ok := true
	T1 := [3]string{"o", "x", "x"}
	T2 := [3]string{"x", "o", "x"}
	T3 := [3]string{"x", "x", "o"}
	fmt.Scan(&S)
	arr := strings.Split(S, "")
	for i := 0; i < len(S); i++ {
		if T1[i%3] != arr[i] {
			T1ok = false
		}

		if T2[i%3] != arr[i] {
			T2ok = false
		}

		if T3[i%3] != arr[i] {
			T3ok = false
		}
	}
	if T1ok || T2ok || T3ok {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}
