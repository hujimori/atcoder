package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INF = 1000010000

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		a, b, t := nextInt()-1, nextInt()-1, nextInt()
		d[a][b] = t
		d[b][a] = t
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	ans := INF
	for i := 0; i < n; i++ {
		tmp := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			tmp = max(tmp, d[i][j])
		}
		ans = min(ans, tmp)
	}

	fmt.Println(ans)

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
