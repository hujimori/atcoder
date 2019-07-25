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

func score(scoreX, scoreY int) int {
	return abs(scoreX - scoreY)
}

const INF = 1000001000

// プレイヤーxがy枚目を引いた時のスコア
var memo [][]int

func dfs(scoreX, scoreY, n int, player int) int {

	if n == N {
		// fmt.Println(score(scoreX, scoreY))
		return score(scoreX, scoreY)
	}

	if memo[player][n] != -1 {
		return memo[player][n]
	}

	if player == 1 {
		ma := -1 * INF
		for i := n; i < N; i++ {
			ma = max(ma, dfs(a[i], scoreY, i+1, 0))
		}
		memo[player][n] = ma
		return memo[player][n]
	} else {
		mi := INF
		for i := n; i < N; i++ {
			mi = min(mi, dfs(scoreX, a[i], i+1, 1))
		}
		memo[player][n] = mi
		return memo[player][n]
	}

	// var res int

	// for i := n; i < N; i++ {
	// 	if player == 1 {
	// 		res = max(-1*INF, dfs(a[i], scoreY, i+1, 0))
	// 		memo[player][i] = res
	// 	} else {
	// 		res = min(INF, dfs(scoreX, a[i], i+1, 1))
	// 		memo[player][i] = res
	// 	}

	// }

	// ans := -1*INF
	// for i := 0; i

	// return memo[player][n]
}

var N, Z, W int
var a []int

func main() {
	sc.Split(bufio.ScanWords)
	N, Z, W = nextInt(), nextInt(), nextInt()

	memo = make([][]int, 2)
	for i := 0; i < 2; i++ {
		memo[i] = make([]int, 2020)
		for j := 0; j < 2020; j++ {
			memo[i][j] = -1
		}
	}

	a = make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}

	fmt.Println(dfs(Z, W, 0, 1))

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
