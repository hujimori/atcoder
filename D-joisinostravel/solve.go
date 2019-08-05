package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

const INF = 1000010000

var used []bool

func main() {
	sc.Split(bufio.ScanWords)

	N, M, R := nextInt(), nextInt(), nextInt()
	r := make([]int, R)
	for i := 0; i < R; i++ {
		r[i] = nextInt() - 1
	}
	sort.Ints(r)

	d := make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([]int, N)
		for j := 0; j < N; j++ {
			d[i][j] = INF
		}
		d[i][i] = 0

	}

	for i := 0; i < M; i++ {
		a, b, c := nextInt()-1, nextInt()-1, nextInt()
		d[a][b] = c
		d[b][a] = c
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	ans := INF
	for {
		path, sum := make([]int, len(r)), 0
		copy(path, r)
		for i := 0; i < len(path)-1; i++ {
			u, v := path[i], path[i+1]
			sum += d[u][v]
		}
		ans = min(ans, sum)
		if !nextPermutation(sort.IntSlice(r)) {
			break
		}
	}
	fmt.Println(ans)

}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}

	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}

	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
