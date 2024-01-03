package main

import "fmt"

func main(){
	func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
		c := make(chan int)
		go func(){
			defer close(c)
			select{
				case a := <- firstChan:
				c <- (a * a)
				case a1 := <- secondChan:
				c <- (a1 * 3)
				case <- stopChan:
				break
			}
		}()
		return c
	}
}