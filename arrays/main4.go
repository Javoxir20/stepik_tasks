/*Given an array consisting of integers. The numbering of elements starts from 0. Write a program that will print the elements of an array whose indices are even (0, 2, 4...).

Input data

Number given first 
�
N— number of elements in the array (
1
≤
�
≤
100
_
_
1≤N≤100). Next, separated by a space, written 
�
N numbers - array elements. The array consists of integers.

Output

It is necessary to display all array elements with even indices.*/

package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i])
			fmt.Print(" ")
		}
	}
}
