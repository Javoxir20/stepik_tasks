package main
import "fmt"

func main() {
    // put your code here
    var a string
    fmt.Scan(&a)
    b := []rune(a)
    d := true
    for i := 0; i < len(b); i++{
        for j := 0; j < len(b); j++{
            if string(b[i]) == string(b[j]) && i != j{
                d = false
    
            }
        }
  if d == true{
   fmt.Print(string(b[i]))
  }
  d = true
        
    }
}