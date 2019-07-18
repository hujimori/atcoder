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

var memo [110][350]int
var ng map[int]int
var n int

func main() {
	sc.Split(bufio.ScanWords)

	n = nextInt()

	ng = map[int]int{}
	for i := 0; i < 3; i++ {
		ng[nextInt()] = 1
	}

	for i := 0; i < 110; i++ {
		for j := 0; j < 350; j++ {
			memo[i][j] = -1
		}
	}

	if dfs(0, n) == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

// iまででsumを作って、残りi移行を調べる
var minus = []int{1, 2, 3}

func dfs(i, sum int) int {

	if memo[i][sum] != -1 {
		return memo[i][sum]
	}

	if _, exist := ng[sum]; exist {
		return 0
	}

	if sum == 0 {
		return 1
	}

	if i >= 100 {
		if sum == 0 {
			return 1
		}
		return 0
	}

	res := 0
	for j := 0; j < len(minus); j++ {
		if sum-minus[j] >= 0 {
			memo[i][sum] = dfs(i+1, sum-minus[j])
			res = res | memo[i][sum]
		}
	}

	return res
}
