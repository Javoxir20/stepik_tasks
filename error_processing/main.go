package main

import "fmt"

func main() {
    // package main уже объявлен, нужные импорты уже есть
    var a, b int
    fmt.Scan(&a,&b)
    if a ==0 || b==0{
        fmt.Println("ошибка")
    }else{
        fmt.Println(a/b)
    }
}
