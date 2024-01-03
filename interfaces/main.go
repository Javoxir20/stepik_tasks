package main

import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	"os"            // пакет используется для проверки ответа, не удаляйте его
)



func main() {
	v1, v2, op := readTask() // исходные данные получаются с помощью этой функции
                                            // все полученные значения имеют тип пустого интерфейса
    if _, ok := v1.(float64); !ok{
        fmt.Printf("value=%v: %T\n", v1, v1)
        return 
    }
    
    if _, ok := v2.(float64); !ok{
        fmt.Printf("value=%v: %T\n", v2, v2)
        return
    }
    v1 = v1.(float64)
    v2 = v2.(float64)
    
    switch op{
        case "+":
        fmt.Printf("%.4f", v1.(float64) + v2.(float64))
        case "-":
        fmt.Printf("%.4f", v1.(float64) - v2.(float64))
        case "*":
        fmt.Printf("%.4f", v1.(float64) * v2.(float64))
        case "/":
        fmt.Printf("%.4f", v1.(float64) / v2.(float64))
        default:
        fmt.Println("неизвестная операция")
        return
    }
}