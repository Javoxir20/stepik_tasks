package  main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a,&b)
	test(&a,&b)
}

func test(x1 *int, x2 *int) {
	// здесь ваш код
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}