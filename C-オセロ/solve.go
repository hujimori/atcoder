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

	N, Q := nextInt(), nextInt()
	l := make([]int, Q)
	r := make([]int, Q)
	for i := 0; i < Q; i++ {
		l[i], r[i] = nextInt(), nextInt()
	}

	field := make([]int, N+1)
	for i := 0; i < Q; i++ {
		field[l[i]-1]++
		field[r[i]]--
	}
	ans := 0
	for i := 0; i < N; i++ {
		ans += field[i]
		fmt.Printf("%d", ans%2)
	}
	fmt.Println()

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
