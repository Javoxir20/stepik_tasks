package main

import "fmt"

func main(){
	sumInt(1,2,3)
}
func sumInt(a ...int) (int,int){
	var cnt, sm int
	for _, i := range a {
		sm += i
		cnt += 1
	}
	return cnt, sm
}	
