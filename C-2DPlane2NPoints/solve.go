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

func (abs ABS) Len() int {
	return len(abs)
}

func (abs ABS) Swap(i, j int) {
	abs[i], abs[j] = abs[j], abs[i]
}

func (abs ABS) Less(i, j int) bool {
	return abs[i].a < abs[j].a
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	ab := make([]AB, n)
	cd := make([]AB, n)

	for i := 0; i < n; i++ {
		ab[i].a, ab[i].b = nextInt(), nextInt()
	}

	for i := 0; i < n; i++ {
		cd[i].a, cd[i].b = nextInt(), nextInt()
	}

	// sort.Sort(ab)
	// sort.Sort(ByB{ab})

	sort.Sort(ABS(ab))
	sort.Sort(ABS(cd))

	// fmt.Println(ab)
	// fmt.Println(cd)
	used := make([]bool, len(cd))
	ans := 0
	for i := 0; i < n; i++ {
		B := cd[i]
		index := -1
		maxv := -1
		for j := 0; j < n; j++ {
			if !used[j] && ab[j].a < B.a && ab[j].b < B.b && maxv < ab[j].b {
				index = j
				maxv = max(maxv, ab[j].b)
			}
		}
		if index != -1 {
			ans++
			used[index] = true
		}
	}
	fmt.Println(ans)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
