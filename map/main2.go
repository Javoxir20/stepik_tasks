package main 

import "fmt"

func main(){
	for i, _ := range cityPopulation {
		isInMap := false
		for _, v := range groupCity[100] {
			if i == v{
				isInMap = true
			}
		}
		if isInMap == false {
			delete(cityPopulation, i)
		}
	}
	
}