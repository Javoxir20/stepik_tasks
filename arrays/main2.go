/*Write a program that takes a number as input 
�
(
�
≥
4
)
N ( N≥4 ), and then 
�
N integer elements of the slice. The output should be the 4th (3rd by index) element of this slice.*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if i == 3 {
			fmt.Println(b)
			break
		}
	}
}