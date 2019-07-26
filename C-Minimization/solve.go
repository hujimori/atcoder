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

var N, K int
var A []int

func main() {
	sc.Split(bufio.ScanWords)
	N, K = nextInt(), nextInt()
	A = make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	res := 0
	right := 0

	for {
		if res == 0 {
			right += K
		} else {
			right += K - 1
		}
		res++

		if right >= N {
			break
		}
	}

	fmt.Println(res)

}
