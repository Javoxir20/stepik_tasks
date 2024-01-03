/*The input is five integers, which are written to an array. 
However, this part of the program has already been written.
You need to write a piece of code that can be used to find and print the maximum number in this array.*/

package main

import "fmt"

func main() {
	array := [5]int{}
	var a, mx int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	mx = array[0]
	for i := 0; i < 5; i++ {
		if array[i] > mx {
			mx = array[i]
		}
	}
	fmt.Println(mx)
}