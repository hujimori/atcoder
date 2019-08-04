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

	N, K := nextInt(), nextInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}
	S := make([]int, N+1)

	for i := 0; i < N; i++ {
		S[i+1] = S[i] + a[i]
	}

	ans := 0
	for i := 0; i+K < N+1; i++ {
		ans += S[i+K] - S[i]
	}

	fmt.Println(ans)
}
