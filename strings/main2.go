package main

import "fmt"

func main(){
 a := ""
 fmt.Scan(&a)
    
 b := []rune(a)
 c := true
 for i := 0; i < len(b) / 2; i ++{
  if b[i] != b[len(b) - 1 - i]{
   c = false
  }
 }
 if c != false{
  fmt.Print("Палиндром")
 }else {
  fmt.Print("Нет")
 }
}