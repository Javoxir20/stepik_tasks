package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	a, b := 1, 1
	for i := 1; a <= n; i++ {
		a, b = b, a+b
		cnt += 1
		if b > n {
			break
		}
	}
	if a == n {
		fmt.Println(cnt + 1)
	} else {
		fmt.Println(-1)
	}
}
