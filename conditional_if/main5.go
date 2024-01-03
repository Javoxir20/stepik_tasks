/*It is required to determine whether a given year is a leap year, recall:
A year is a leap year if it meets at least one of the following conditions:
- is a multiple of 400;
- is a multiple of 4, but not a multiple of 100.

Input data

A single number is entered - the year number (integer, positive, does not exceed 10000).

Output

You want to output the word YES if the year is a leap year and NO otherwise.*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%400 == 0 || n%4 == 0 && n%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}