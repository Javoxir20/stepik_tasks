/*You need to write a program that, when executed, reads two natural numbers A and B from the keyboard 
(each no more than 100, A < B). Print the sum of all numbers from A to B inclusive.*/

package main

import "fmt"

func main() {
	var a, b, sm int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		sm += i
	}
	fmt.Println(sm)
}