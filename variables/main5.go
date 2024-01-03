/*Given a non-negative integer. Find the tens number (that is, the second digit from the right). 

Input data format
The input is a natural number not exceeding 10000.

Output format
Print one integer - the number of tens.

*/

package main

import "fmt"

func main(){
    var n int
    fmt.Scan(&n)
    fmt.Println(n / 10 % 10)
}