/*Find  the first  number from 1 to n inclusive that is a multiple of c but  NOT  a multiple of d.

Input data

Enter 3 natural numbers n, c, d, each of which does not exceed 10000.

Output
Print the first number from 1 to n inclusive, a multiple of c, but  NOT  a 
multiple of d. If there is no such number, you do not need to output anything.*/

package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}