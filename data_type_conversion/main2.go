package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var a, b string
	fmt.Scan(&a,&b)
	adding(a,b)
}
func adding(a, b string) int64{
	a1:=[]rune(a)
	b1:=[]rune(b)
    var cnt, sm string
    for _,i := range a1{
		if unicode.IsDigit(i){
			cnt += string(i)
		}
	}
	for _,i := range b1{
		if unicode.IsDigit(i){
			sm += string(i)
		}
	}
	ans1,_ := strconv.Atoi(cnt)
	ans2,_ := strconv.Atoi(sm)
    return  int64(ans1 + ans2)
}