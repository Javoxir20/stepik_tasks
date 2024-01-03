package main

import "fmt"

func main() {
	var n, a, cnt int
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}
