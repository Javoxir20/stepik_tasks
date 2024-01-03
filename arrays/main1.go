/*Inside the main function (there is no need to declare the function), you need to write a program:

In the first step, 10 positive integers are supplied to standard input 
and must be written in order of input into a 10-element array. The type of numbers 
included in the array must match the smallest possible unsigned integer . The name of 
the array that you must create yourself  workArray (required condition) . To read from 
standard input,  the fmt package is already imported .

At the second stage, 3 more pairs of numbers are supplied to the standard input - the indices of 
the elements of this array that need to be swapped (if such a pair of numbers is 3 and 7, then in 
	the array the element with index 3 must be swapped with the element whose index is 7).

The elements of the resulting array must be printed, separated by spaces, to standard output. Next, 
the types used will be automatically checked, the result of which will be added to your response.*/

package main

import "fmt"

func main(){
	var workArray [10]uint8
for i := 0; i < 10; i++ {
	fmt.Scan(&workArray[i])
}
var a, b int
for i := 3; i > 0; i-- {
	fmt.Scan(&a, &b)
	workArray[a], workArray[b] = workArray[b], workArray[a]
}
for _, i := range workArray {
	fmt.Print(i, " ")
}

}