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

const MAX = 1000
const INF = 1000010000

func main() {
	sc.Split(bufio.ScanWords)

	H, W := nextInt(), nextInt()

	c := make([][]int, 10)
	for i := 0; i < 10; i++ {
		c[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			c[i][j] = nextInt()
		}
	}

	A := make([][]int, H)
	for i := 0; i < H; i++ {
		A[i] = make([]int, W)
		for j := 0; j < W; j++ {
			A[i][j] = nextInt()
		}
	}

	for k := 0; k < 10; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				c[i][j] = min(c[i][j], c[i][k]+c[k][j])
			}
		}
	}

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if 0 <= A[i][j] {
				ans += c[A[i][j]][1]
			}
		}
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
