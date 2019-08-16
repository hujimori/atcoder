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
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		A[i], B[i] = nextInt(), nextInt()
	}

	p := make([]int, N+1)
	for i := N - 1; i >= 0; i-- {
		A[i] += p[i+1]

		amari := A[i] % B[i]
		plus := 0
		if amari != 0 {
			plus = B[i] - amari
		}
		p[i] = p[i+1] + plus
	}
	fmt.Println(p[0])
}
