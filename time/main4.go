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
	b := strings.Replace(a, " мин. ", "m", 1)
	b = strings.Replace(b, " сек.", "s", 1)
	b = strings.Replace(b, "\n", "", 1)
	c, err := time.ParseDuration(b)
	if err != nil {
		panic(err)
	}
	const now = 1589570165
	ut := time.Unix(now, 0).UTC()
	ans := ut.Add(c)
	fmt.Println(ans.Format(time.UnixDate))
}