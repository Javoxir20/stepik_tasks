package main

import "fmt"

func main() {
 var a string
 fmt.Scan(&a)
 b := []rune(a)
 for i := 0; i < len(b); i++ {
  switch string(b[i]) {
  case "0":
   fmt.Print("0")

  case "9":
   fmt.Print("81")

  case "1":
   fmt.Print("1")

  case "2":
   fmt.Print("4")

  case "3":
   fmt.Print("9")

  case "4":
   fmt.Print("16")

  case "5":
   fmt.Print("25")

  case "6":
   fmt.Print("36")

  case "7":
   fmt.Print("49")

  case "8":
   fmt.Print("64")
  }
 }
}