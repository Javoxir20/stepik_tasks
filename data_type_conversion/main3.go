package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	s := strings.ReplaceAll(input, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	s1 := strings.Split(s, ";")
	n1, _ := strconv.ParseFloat(s1[0], 64)
	n2, _ := strconv.ParseFloat(s1[1], 64)
	fmt.Printf("%.4f", n1/n2)
}