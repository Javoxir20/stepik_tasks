package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
    //print your code here
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
