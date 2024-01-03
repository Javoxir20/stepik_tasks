package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	b := strings.Split(a, ",")
	ft, err := time.Parse("02.01.2006 15:04:05", b[0][:19])
	if err != nil{
		panic(err)
	}
	st, err := time.Parse("02.01.2006 15:04:05", b[1][:19])
	if err != nil {
		panic(err)
	}
	if ft.After(st){
		fmt.Println(ft.Sub(st))
	}else{
		fmt.Println(st.Sub(ft))
	}
}