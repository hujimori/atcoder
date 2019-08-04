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

	N, M := nextInt(), nextInt()

	ans := 0
	if N <= M/2 {
		ans = (N + M/2) / 2
	} else {
		ans = M / 2
	}

	fmt.Println(ans)

}
