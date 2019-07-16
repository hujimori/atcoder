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

var h int
var w int
var field [][]int
var dp [1100][1100]float64

var dx = []int{1, 0, -1, 0, 0}
var dy = []int{0, 1, 0, -1, 0}

const mod = 1000000007

func rec(x, y int) int {
	if dp[x][y] != -1 {
		return int(dp[x][y])
	}

	res := 0
	for i := 0; i < len(dx); i++ {
		sx := x + dx[i]
		sy := y + dy[i]

		if (0 <= sx && sx < h) && (0 <= sy && sy < w) && field[x][y] < field[sx][sy] {
			res += rec(sx, sy)
			res %= mod
		}
	}
	dp[x][y] = float64(res%mod + 1)
	return int(dp[x][y])

}

func main() {
	sc.Split(bufio.ScanWords)

	h, w = nextInt(), nextInt()

	field = make([][]int, h)
	for i := 0; i < h; i++ {
		field[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			field[i][j] = nextInt()
			dp[i][j] = -1
		}
	}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans += rec(i, j)
			ans %= mod
		}
	}
	fmt.Println(ans)
}

func chmin(a *float64, b float64) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax(a *float64, b float64) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
