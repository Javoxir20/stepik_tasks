/*The hour hand has turned d degrees since the beginning of the day. 
Determine how many whole hours h and whole minutes m there are now.

Input data

The program input is an integer d (0 <  d  < 360).

Output

Display the phrase:

It is ... hours ... minutes.

Instead of ellipses, the program should print the values ​​of h 
and m, separating them from the words with exactly one space.*/



package main

import "fmt"

func main() {
    var d, h, m int
    fmt.Scan(&d)
    h = d * 2 / 60
    m = d % 30 * 2
    fmt.Println("It is", h, "hours", m, "minutes.")
}