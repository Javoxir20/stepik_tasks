/*Given a  three-digit  number, determine whether all its digits  are different .

Input data format
The input is one natural three-digit number.

Output format
Print "YES" if all digits of the number are different, otherwise print "NO".*/

package main

import "fmt"

func main() {
	var n int
	var y, u, b int
	fmt.Scan(&n)
	y = n / 100
	u = n / 10 % 10
	b = n % 10
	if y != u && y != b && u != b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}