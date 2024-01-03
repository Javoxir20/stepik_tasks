package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := 1
	for i := 0; a <= n; i++ {
		fmt.Print(a, " ")
		a *= 2
	}
}
