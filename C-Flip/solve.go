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

	n, m := uint64(nextInt()), uint64(nextInt())
	if n == 1 && m == 1 {
		fmt.Println(1)
		return
	} else if n == 1 || m == 1 {
		fmt.Println(max(int(n), int(m)) - 2)
		return
	}

	outer := (n-2)*2 + m*2
	inner := n*m - outer
	fmt.Println(inner)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a

}
