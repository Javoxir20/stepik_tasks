package main 
import "fmt"

func main(){
    var b string
    fmt.Scan(&b)
 a := []rune(b)
    for i := 0; i < len(a); i++{
        if i % 2 == 1{
   fmt.Print(string(a[i]))
  }
    }
}