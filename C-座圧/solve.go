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

func main() {
	sc.Split(bufio.ScanWords)

	N := nextInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}

	m := make(map[int]int)
	for _, x := range a {
		m[x] = 0
	}
	u := make([]int, 0, len(m))
	for x := range m {
		u = append(u, x)
	}
	sort.Ints(u)
	for i := 0; i < len(u); i++ {
		m[u[i]] = i
	}
	for _, x := range a {
		fmt.Println(m[x])
	}

	// aa := make([]int, N)
	// copy(aa, a)

	// sort.Ints(aa)
	// m := map[int]int{}
	// for i := 0; i < N; i++ {
	// 	m[aa[i]] = i
	// }

	// for i := 0; i < N; i++ {
	// 	fmt.Println(m[a[i]])
	// }
}
