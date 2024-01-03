package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	mx := 0
	cnt := 0
	for i := a; i <= b; i++ {
		if i%7 == 0 {
			mx = i
		} else {
			cnt += 1
		}
	}
	// fmt.Print(cnt)
	if cnt == b-a+1 {
		fmt.Print("NO")
	} else {
		fmt.Print(mx)
	}
}