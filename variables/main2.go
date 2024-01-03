/*Petya was in a hurry to go to school and incorrectly wrote 
a program that first finds the squares of two numbers and then sums them. Fix his program.*/
package main

import "fmt"

func main(){

  var a, b, c int
    
  fmt.Scan(&a) // считаем переменную 'a' с консоли
  fmt.Scan(&b) // считаем переменную 'b' с консоли

  a *= a
  b *= b 
  c = a + b
  fmt.Println(c)
}