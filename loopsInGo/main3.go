/*Write a program that, in a sequence of numbers, finds the 
sum of two-digit numbers that are multiples of 8. The program in 
the first line receives as input the number n - the number of numbers 
in the sequence, in the second line - n numbers included in this sequence.*/

package main

import "fmt"

func main() {
	var n, a, sm int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a < 99 && a > 10 && a%8 == 0 {
			sm += a
		}
	}
	fmt.Println(sm)
}
