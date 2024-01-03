/*Two numbers are given. Determine the digits included in the entry of both the first and second numbers.

Input data

The program receives two numbers as input. It is guaranteed that digits 
in numbers are not repeated. Numbers ranging from 0 to 10000.

Output

The program should output the digits that are present in 
both numbers, separated by a space. The numbers are displayed in the order they appear in the first number.

Attention ! If you find it difficult to solve this problem, 
skip and study the material further, then return to this problem later.*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var num, num2 int
	for a != 0 {
		num = num*10 + (a % 10)
		a = a / 10

	}
	for b != 0 {
		num2 = num2*10 + (b % 10)
		b = b / 10

	}
	a = num
	b = num2

	for a != 0 {
		c := a % 10
		var cnt int
		d := b

		for d != 0 {
			i := d % 10

			if c == i {
				cnt++
			}

			d = d / 10

		}

		if cnt > 0 {
			fmt.Print(c, " ")
		}

		a = a / 10
	}

}