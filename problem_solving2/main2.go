package  main
import "fmt"

func main(){
    var a string
    fmt.Scan(&a)
    b := []rune(a)
    for i := 0; i < len(a) - 1; i++ {
        fmt.Print(string(b[i]),"*")
    }
    fmt.Print(string(b[len(b) - 1]))
}