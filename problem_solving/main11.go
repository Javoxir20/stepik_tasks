package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)
	if n > 5 && n < 20 {
		fmt.Printf("%d korov", n)
	} else {
		a = n % 10
		if a == 1 {
			fmt.Printf("%d korova", n)
		} else if a == 2 || a == 3 || a == 4 {
			fmt.Printf("%d korovy", n)
		} else {
			fmt.Printf("%d korov", n)
		}
	}
}