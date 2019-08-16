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
	sc.Split(bufio.ScanWords)

	A, B, C := nextInt(), nextInt(), nextInt()
	for i := 0; i < 1000; i++ {
		if A%2 == 1 {
			fmt.Println(i)
			return
		}
		if B%2 == 1 {
			fmt.Println(i)
			return
		}
		if C%2 == 1 {
			fmt.Println(i)
			return
		}

		ta := (B + C) / 2
		tb := (A + C) / 2
		tc := (A + B) / 2

		A = ta
		B = tb
		C = tc

	}

	fmt.Println(-1)
}
