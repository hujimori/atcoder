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

	N := nextInt()

	if N == 0 {
		fmt.Println(0)
		return
	}

	if N < 0 {
		N = -1 * N
	}

	var ans string

	for {
		ans += strconv.Itoa(N % 2)
		N = N / 2

		if N == 0 {
			break
		}
	}

	fmt.Println(ans)
}
