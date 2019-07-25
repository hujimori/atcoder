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

var G [][]int
var N int
var dpF []int
var dpG []int

const MOD = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	G = make([][]int, N)
	for i := 0; i < N; i++ {
		G[i] = make([]int, 0)
	}

	for i := 0; i < N-1; i++ {
		u, v := nextInt()-1, nextInt()-1
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	dpF = make([]int, N)
	dpG = make([]int, N)

	ans := f(0, -1) % MOD
	fmt.Println(ans)

}

func f(x, p int) int {
	if dpF[x] > 0 {
		return dpF[x]
	}

	white := g(x, p)
	black := 1
	for _, child := range G[x] {
		if child == p {
			continue
		}

		black *= g(child, x) % MOD
		black %= MOD
	}
	dpF[x] = (white + black) % MOD
	return dpF[x]
}

func g(x, p int) int {

	if dpG[x] > 0 {
		return dpG[x]
	}

	count := 1
	for _, child := range G[x] {
		if child == p {
			continue
		}

		count *= f(child, x) % MOD
		count %= MOD
	}
	dpG[x] = count
	return dpG[x]

}
