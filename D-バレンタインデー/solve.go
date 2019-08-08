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

type edge struct {
	to   int
	cost int
}

func nextCombination(sub int) int {
	x := sub & -sub
	y := sub + x
	return (((sub & ^y) / x) >> 1) | y
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M, P, Q, R := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	g := make([][]edge, N)

	for i := 0; i < R; i++ {
		x, y, z := nextInt()-1, nextInt()-1, nextInt()
		g[x] = append(g[x], edge{to: y, cost: z})
	}

	ans := -1
	bit := (1<<uint(P) - 1)
	for ; bit < (1 << uint(N)); bit = nextCombination(bit) {
		value := make([]int, M)

		for i := uint(0); i < uint(N); i++ {
			if (bit & (1 << i)) == 1<<i {
				for j := 0; j < len(g[i]); j++ {
					to := g[i][j].to
					cost := g[i][j].cost
					value[to] += cost
				}
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(value)))
		sum := 0
		for i := 0; i < Q; i++ {
			sum += value[i]
		}
		ans = max(ans, sum)

	}

	fmt.Println(ans)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
