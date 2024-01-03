/*Write a program that sequentially performs the following operations with the entered number:

The number is multiplied by 2;
Then 100 is added to the number.

After this, the resulting number should be displayed on the screen.*/
package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)
    // здесь ваш код
    a *= 2
    a += 100
    fmt.Println(a)
}
