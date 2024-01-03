package main

import "fmt"

func main(){
	var n int64
	fmt.Scan(&n)
	convert(n)
}
func convert(a int64) uint16{
    return uint16(a)
}