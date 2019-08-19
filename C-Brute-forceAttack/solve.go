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

func f(rest int, s string) {
	if rest == 0 {
		fmt.Println(s)
	} else {
		for c := 'a'; c <= 'c'; c++ {
			f(rest-1, s+string(c))
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := nextInt()
	f(N, "")

}
