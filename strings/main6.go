package main
import (
 "fmt"
    "bufio"
 "os"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    b := []rune(text) 
    length := len(b)
 c := true
    if length < 5 {
        fmt.Print("Wrong password")
    }else {
        lat_num := 0
        num := 0
        for i := 0; i < length; i++{
            if string(b[i]) >= "A" && string(b[i]) <= "Z" ||string(b[i]) >= "a" && string(b[i]) <= "z" {
    lat_num += 1
            }else if string(b[i]) >= "0" && string(b[i]) <= "9" {
    num += 1
   }else{
    c = false
   }
        }
  if c == false{
   fmt.Print("Wrong password")
  }else if num > 0 || lat_num > 0{
   fmt.Print("Ok")
  } else {
   fmt.Print("Wrong password")
  }
    }

 
}