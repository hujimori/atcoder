package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	var s string
	fmt.Scan(&s)

	slice := strings.Split(s, "")

	sum := len(s)
	bCnt := 0
	for i := 0; i < len(s); i++ {
		if slice[i] == "B" {
			continue
		}
		if slice[i] 
		

	}
}
