/*The bank deposit is  x  rubles. Every year it increases by  p  percent, after which the fractional part of the kopecks is discarded. Every year the deposit amount becomes larger. Determine how many years later the deposit will be at least  y  rubles.

Input data

The program receives three natural numbers as input:  x ,  p ,  y .

Output

The program should output one integer.*/
package main

import "fmt"

func main() {
	var x, p, y, ans, cnt int
	fmt.Scan(&x, &p, &y)
	for x <= y {
		ans = x * p / 100
		x += ans
		cnt += 1
	}
	fmt.Println(cnt)

}
