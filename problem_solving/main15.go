package main

import (
	"fmt"
)

func main() {
	var a, b int
	rv := 0
	fmt.Scan(&a, &b)
	n := a
	var rv1 int
    for a > 0{
		rv1 = rv1 *10 + a % 10
		a /= 10
	}
	for rv1 > 0 {
		if rv1%10 != b{
			rv = rv*10 + rv1%10
		}
		rv1 /= 10
	}
	
	if n % 10 == 0{
	    rv *= 10
	}
	fmt.Println(rv)
}