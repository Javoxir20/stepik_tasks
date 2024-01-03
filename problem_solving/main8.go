package main

import "fmt"

func main() {
	var sls []int
	var n, a, mn int
	var cnt int = 1
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sls = append(sls, a)
		mn = sls[0]
	}
	for i := 1; i < n; i++ {
		if sls[i] < mn {
			mn = sls[i]
			cnt = 1
		} else if mn == sls[i] {
			cnt += 1
		}
	}
	fmt.Println(cnt)

}
