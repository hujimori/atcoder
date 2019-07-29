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

type Point struct {
	pos int
	idx int
}

type ByPos []Point

func (b ByPos) Len() int {
	return len(b)
}

func (b ByPos) Less(i, j int) bool {
	return b[i].pos < b[j].pos
}

func (b ByPos) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type Adjacent struct {
	distance int
	a        int
	b        int
}

type ByDistance []Adjacent

func (b ByDistance) Len() int {
	return len(b)
}

func (b ByDistance) Less(i, j int) bool {
	return b[i].distance < b[j].distance
}

func (b ByDistance) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	sc.Split(bufio.ScanWords)

	N := nextInt()
	xP := make([]Point, N)
	yP := make([]Point, N)
	ut := NewUnionFindTree(N)

	for i := 0; i < N; i++ {
		xP[i] = Point{pos: nextInt(), idx: i}
		yP[i] = Point{pos: nextInt(), idx: i}
	}
	sort.Sort(ByPos(xP))
	sort.Sort(ByPos(yP))

	adj := make([]Adjacent, (N-1)*2)
	for i := 0; i < N-1; i++ {
		adj[2*i] = Adjacent{
			distance: xP[i+1].pos - xP[i].pos,
			a:        xP[i].idx,
			b:        xP[i+1].idx,
		}

		adj[2*i+1] = Adjacent{
			distance: yP[i+1].pos - yP[i].pos,
			a:        yP[i].idx,
			b:        yP[i+1].idx,
		}
	}
	sort.Sort(ByDistance(adj))

	ans := 0
	for i := 0; i < (N-1)*2; i++ {
		if !ut.Same(adj[i].a, adj[i].b) {
			ans += adj[i].distance
			ut.Unit(adj[i].a, adj[i].b)
		}
	}

	fmt.Println(ans)

}
