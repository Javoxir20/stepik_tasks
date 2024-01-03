/*Given a non-negative integer. Find and print the first digit of the number. 

Input data format
The input is a natural number not exceeding 10000.

Output format
Print one integer - the first digit of the given number.*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for !(n >= 1 && n <= 9) {
		n = n / 10
	}
	fmt.Println(n)
}
