package main

import "fmt"

func main() {
	var n, ans, sm int
	fmt.Scan(&n)

	for n != 0 {
		sm += n % 10
		n = n / 10
	}
	for sm != 0 {
		ans += sm % 10
		sm = sm / 10
	}
	fmt.Println(ans)
}
