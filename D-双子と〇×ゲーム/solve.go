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

func score(board [][]int) int {
	ans := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == board[i+1][j] {
				ans += b[i][j]
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if board[i][j] == board[i][j+1] {
				ans += c[i][j]
			}
		}
	}

	return ans
}

const INF = 1000001000

func dfs(board [][]int, u int) int {
	if u == 9 {
		return score(board)
	}

	var minmax int
	if u%2 == 0 {
		minmax = -1 * INF
	} else {
		minmax = INF
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != 0 {
				continue
			}
			if u%2 == 0 {
				board[i][j] = 1
				minmax = max(minmax, dfs(board, u+1))
				board[i][j] = 0
			} else {
				board[i][j] = 2
				minmax = min(minmax, dfs(board, u+1))
				board[i][j] = 0
			}
		}
	}
	return minmax
}

var b [2][3]int
var c [3][2]int
var sum int
var board [][]int

func main() {
	sc.Split(bufio.ScanWords)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = nextInt()
			sum += b[i][j]
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = nextInt()
			sum += c[i][j]
		}
	}

	board = make([][]int, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]int, 3)
	}

	cho := dfs(board, 0)
	fmt.Println(cho)
	fmt.Println(sum - cho)

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
