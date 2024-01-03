/*The input is a  float64 number . You need to output the 
converted number according to the rule: the number is squared, 
then the fractional part is cut off so that 4 decimal places remain. 
But if the number is equal to or less than zero, output:
“the number R is not suitable,” where R is the original number with 
2 digits after the decimal point and 2 in width. And if the number is more 
than 10,000, display the original number in exponential notation (the exponent sign is in lowercase).*/

package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	if num > 10000 {
		fmt.Printf("%e\n", num)
	} else if num <= 0 {
		fmt.Printf("число %2.2f не подходит\n", num)
	} else {
		fmt.Printf("%.4f\n", num*num)
	}
}
