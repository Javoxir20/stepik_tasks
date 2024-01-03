/*Write a program that reads integers from the console, one number per line.

For each number entered, check:

if the number is less 10, then we skip this number;
if the number is greater 100, then we stop reading the numbers;
in other cases, print this number back to the console in a separate line.*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 100 {
		if n > 9 {
			fmt.Println(n)
			main()
		} else if n < 10 {
			main()
		}
	}

}