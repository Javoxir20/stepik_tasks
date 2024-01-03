package main

import "fmt"

func main(){
	func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
		c := make(chan int)
		sm := 0
		go func(){
			defer close(c)
			for{
				select{
				case a := <- arguments:
				sm += a
				case <- done:
				c <- sm
				return 
			}
			}
		}()
		return c
	}
}