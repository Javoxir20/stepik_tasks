/*Given a natural number, print its last digit.

Input data format
The input is a natural number N, not exceeding 10000.

Output format
Print one integer - the answer to the problem.*/

package main

import "fmt"

func main(){
    var n int
    fmt.Scan(&n)
    fmt.Println(n % 10)
}