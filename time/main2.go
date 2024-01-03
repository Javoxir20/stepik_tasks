package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	b, err := time.Parse(time.DateTime, a[:19])
	if err != nil {
		panic(err)
	}
	if b.Hour() > 13{
		c := b.AddDate(0, 0, 1)
		fmt.Println(c.Format(time.DateTime))
	}else{
		fmt.Print(a)
	}
}