package main

import (
	"bufio"
	"os"
	"strconv"
)

type UnionFindTree struct {
	par   []int // 親
	rank  []int // 木の深さ
	count int   // 要素数
}

func NewUnionFindTree(count int) *UnionFindTree {
	par := make([]int, count)
	rank := make([]int, count)
	for i := 0; i < count; i++ {
		par[i] = i
		rank[i] = 0
	}
	return &UnionFindTree{par: par, rank: rank, count: count}
}

func (uf UnionFindTree) Find(x int) int {
	if uf.par[x] == x {
		return x
	} else {
		uf.par[x] = uf.par[uf.par[x]]
		return uf.Find(uf.par[x])
	}
}

func (uf *UnionFindTree) Unit(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	} else if uf.rank[x] < uf.rank[y] {
		uf.par[x] = y
	} else {
		uf.par[y] = x
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func (uf UnionFindTree) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

var (
	N int
	K int
	T []int
	X []int
	Y []int
)

func main() {
	sc.Split(bufio.ScanWords)

	N, K = nextInt(), nextInt()

	// Union-Find木を初期化
	// x, x + N, x+2*N　を x-A, x-B, x-Cの要素とする
	u := NewUnionFindTree(3 * N)
	T = make([]int, K)
	X = make([]int, K)
	Y = make([]int, K)
	for i := 0; i < K; i++ {
		T[i], X[i], Y[i] = nextInt(), nextInt(), nextInt()
	}

	ans := 0
	for i := 0; i < K; i++ {
		t := T[i]
		x, y := X[i]-1, Y[i]-1

		// 正しくない番号
		if x < 0 || N <= x || y < 0 || N < y {
			ans++
			continue
		}

		if t == 1 {
			// 「xとyは同じ種類」という情報
			if u.Same(x, y+N) || u.Same(x, y+2*N) {
				ans++
			} else {
				u.Unit(x, y)
				u.Unit(x+N, y+N)
				u.Unit(x+N*2, y+N*2)
			}
		} else {
			// 「xはyを食べる」という情報
			if u.Same(x, y) || u.Same(x, y+2*N) {
				ans++
			} else {
				u.Unit(x, y+N)
				u.Unit(x+N, y+2*N)
				u.Unit(x+N*2, y)
			}
		}
	}

	println(ans)
}
