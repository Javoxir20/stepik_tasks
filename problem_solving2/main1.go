package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	c = a*a + b*b
	fmt.Println(math.Sqrt(float64(c)))
}
