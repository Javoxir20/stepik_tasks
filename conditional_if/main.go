/*The input is an integer. If the number is positive - display the message 
"The number is positive", if the number 
is negative - "The number is negative". If zero is supplied, display the message “Zero”. Display 
the message without quotes.*/

package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    if n > 0{
        fmt.Println("The number is positive")
    } else if n < 0{
        fmt.Println("The number is negative")
    }else{
        fmt.Println("zero")
    }
}