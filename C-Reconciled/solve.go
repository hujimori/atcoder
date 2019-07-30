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

const MOD = 1000000007

func main() {
	sc.Split(bufio.ScanWords)

	N, M := nextInt(), nextInt()

	if abs(N-M) > 1 {
		fmt.Println(0)
		return
	}

	n := N
	for i := n - 1; i >= 1; i-- {
		n = n * i % MOD
	}

	m := M
	for i := m - 1; i >= 1; i-- {
		m = m * i % MOD
	}

	if n == m {
		fmt.Println(n * m * 2 % MOD)
	} else {
		fmt.Println(n * m % MOD)
	}

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
