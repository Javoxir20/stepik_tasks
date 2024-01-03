package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n/100 + n/10%10 + n%10)
}