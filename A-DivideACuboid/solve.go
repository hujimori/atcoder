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

	ab := A * B
	ac := A * C
	bc := B * C

	if A%2 == 0 || B%2 == 0 || C%2 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(min(ab, min(ac, bc)))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
