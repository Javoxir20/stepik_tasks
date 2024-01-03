package main

import (
	"fmt"
	"time"
)

func main() {
	var a string
	fmt.Scan(&a)
	b, err := time.Parse(time.RFC3339, a)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.Format(time.UnixDate))
}