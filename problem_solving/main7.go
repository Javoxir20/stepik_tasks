package main

import "fmt"

func main() {
	var n, a, cnt int
	var arr []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
		if arr[i] == 0 {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}