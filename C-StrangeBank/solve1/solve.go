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

var (
	n    int
	memo []int
)

var max = 110000

func main() {
	sc.Split(bufio.ScanWords)

	n = nextInt()
	memo = make([]int, max)

	for i := 0; i < max; i++ {
		memo[i] = -1
	}

	fmt.Println(rec(n))
}

func rec(n int) int {
	// 終端条件
	if n == 0 {
		return 0
	}

	// すでに探索済みならリターン
	if memo[n] != -1 {
		return memo[n]
	}

	res := n

	// 1, 6, 6^2, ..., を試す（nを超えない範囲）
	for pow6 := 1; pow6 <= n; pow6 *= 6 {
		res = min(res, rec(n-pow6)+1)
	}

	// 1, 9, 9^2, ..., を試す（nを超えない範囲）
	for pow9 := 1; pow9 <= n; pow9 *= 9 {
		res = min(res, rec(n-pow9)+1)
	}

	// メモしながらリターン
	memo[n] = res
	return memo[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
