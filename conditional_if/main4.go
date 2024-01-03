/* Determine whether the ticket is  lucky . A ticket is considered lucky  
if its six-digit number has the sum of the first  three  digits  equal  to the sum of the last three.
Input format
The ticket number is given at the entrance - one six-digit number.
Output format
Print "YES" if the ticket is lucky, otherwise print "NO". */

package main

import "fmt"

func main() {
	var n, o, b, t, u, i, br int
	fmt.Scan(&n)
	o = n / 100000
	b = n / 10000 % 10
	t = n / 1000 % 10
	u = n / 100 % 10
	i = n / 10 % 10
	br = n % 10
	if o+b+t == u+i+br {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}