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

	var s string
	fmt.Scan(&s)

	slice := strings.Split(s, "")

	cnt := 0

	for i := 0; i < len(s)-1; i++ {
		if slice[i] != slice[i+1] {
			cnt++
		}

	}

	fmt.Println(cnt)
}
