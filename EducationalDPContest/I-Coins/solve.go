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

func nextFloat() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

var n int
var p []float64
var dp [3100][3100]float64

func main() {
	sc.Split(bufio.ScanWords)

	n = nextInt()
	p = make([]float64, n)
	for i := 0; i < n; i++ {
		p[i] = nextFloat()
	}

	// fmt.Println(p)
	dp[0][0] = 1.0

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			dp[i+1][j+1] += dp[i][j] * p[i]
			dp[i+1][j] += dp[i][j] * (1.0 - p[i])
		}
	}

	res := 0.0
	for j := (n + 1) / 2; j <= n; j++ {
		res += dp[n][j]
	}
	fmt.Println(res)

}
