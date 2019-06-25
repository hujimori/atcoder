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

type AB struct {
	a int
	b int
}

type ABS []AB

func (a ABS) Less(i, j int) bool {
	return a[i].a < a[j].a
}

func (a ABS) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ABS) Len() int {
	return len(a)
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()

	ab := make([]AB, n)
	for i := 0; i < n; i++ {
		ab[i].a, ab[i].b = nextInt(), nextInt()
	}

	sort.Sort(ABS(ab))
	count := 0
	ans := 0
	for i := 0; i < n; i++ {
		if count < m {
			c := min(m-count, ab[i].b)
			ans += ab[i].a * c
			count += c

		}
	}
	fmt.Println(ans)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
