package main

import "fmt"

func main(){
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	vote(a,b,c)
}
func vote(x int, y int, z int) int {
    //print your code here
    if x ==y {
        return x
    }else {
        return z
    }    
}