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
	T := make([]int, N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		T[i], A[i] = nextInt(), nextInt()
	}

	t, a := 1, 1

	for i := 0; i < N; i++ {
		x := t / T[i]
		y := a / A[i]
		if t%T[i] != 0 {
			x++
		}

		if a%A[i] != 0 {
			y++
		}
		n := max(x, y)

		t = n * T[i]
		a = n * A[i]

	}

	fmt.Println(t + a)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
