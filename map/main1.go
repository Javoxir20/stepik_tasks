package main

import "fmt"

func main() {
	var a int
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if el, ok := m[a]; ok {
			fmt.Print(el, " ")
		} else {
			m[a] = work(a)
			fmt.Print(m[a], " ")
		}
	}

}
func work(a int) int{
	return a
}
