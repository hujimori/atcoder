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
	return abs[i].b < abs[j].b
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
	sort.Sort(sort.Reverse(ABS(cd)))

	fmt.Println(ab)
	fmt.Println(cd)

	ans := 0
	for i := 0; i < n; i++ {
		if ab[i].a < cd[i].a && ab[i].b < cd[i].b {
			ans++
		}
	}
	fmt.Println(ans)

}
