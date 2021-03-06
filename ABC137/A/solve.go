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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	A, B := nextInt(), nextInt()

	ans := -100000000000

	ans = max(ans, A+B)
	ans = max(ans, A-B)
	ans = max(ans, A*B)
	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
