package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

var h int
var w int
var field [][]string
var dp [1100][1100]int

var dx = []int{1, 0}
var dy = []int{0, 1}

const mod = 1000000007

// func rec(x, y int) int {
// 	if dp[x][y] != -1 {
// 		return int(dp[x][y])
// 	}

// 	res := 0
// 	for i := 0; i < len(dx); i++ {
// 		sx := x + dx[i]
// 		sy := y + dy[i]

// 		if (0 <= sx && sx < h) && (0 <= sy && sy < w) && field[sx][sy] == "." {
// 			res = rec(sx, sy)
// 			dp[x][y] = res%mod + 1
// 		}
// 	}

// 	return int(dp[x][y])
// }

func main() {
	sc.Split(bufio.ScanWords)

	h, w = nextInt(), nextInt()

	field = make([][]string, h)
	for i := 0; i < h; i++ {
		field[i] = make([]string, w)
	}

	for i := 0; i < h; i++ {
		s := strings.Split(nextLine(), "")
		for j := 0; j < w; j++ {
			field[i][j] = s[j]
		}
	}
	// rec(0, 0)
	// fmt.Println(rec(0, 0))

	dp[0][0] = 1

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i+1 < h && field[i+1][j] == "." {
				dp[i+1][j] += (dp[i][j])
				dp[i+1][j] %= mod
			}

			if j+1 < w && field[i][j+1] == "." {
				dp[i][j+1] += dp[i][j]
				dp[i][j+1] %= mod
			}

		}
	}

	fmt.Println(dp[h-1][w-1])
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
