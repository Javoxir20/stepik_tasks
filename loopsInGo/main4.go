/*The sequence consists of natural numbers and ends with the number 0. Determine 
the number of elements of this sequence that are equal to its largest element.

Input format

A non-empty sequence of natural numbers is introduced, ending with the number 0
 (the number 0 itself is not included in the sequence, but serves as a sign of its end).

Output format

Print the answer to the problem.*/

package main

import "fmt"

func main() {
	var n, mx int
	var cnt = 1
	fmt.Scan(&n)
	for n != 0 {
		if n > mx {
			mx = n
			cnt = 1
		} else if mx == n {
			cnt += 1
		}
		fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
