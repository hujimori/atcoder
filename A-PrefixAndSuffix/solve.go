package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	// sc.Split(bufio.ScanWords)

	var N int
	fmt.Scan(&N)
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)


	for l := 0; l < N; l++ {
		if s[l:N] == t[0:N-l] {
			fmt.Println(l + N)
			return
		}
	}

	fmt.Println(N * 2)
}
