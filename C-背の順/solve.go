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

type pair struct {
	fisrt  int
	second int
}

type SortBy []pair

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].second < a[j].second }

func main() {
	sc.Split(bufio.ScanWords)

	N := nextInt()
	p := make([]pair, N)
	for i := 0; i < N; i++ {
		p[i].fisrt = i + 1
		p[i].second = nextInt()
	}

	sort.Sort(sort.Reverse(SortBy(p)))

	for i := 0; i < N; i++ {
		fmt.Println(p[i].fisrt)
	}
}
