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

	n := nextInt()
	a := make([]int, n)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		m[a[i]]++
	}

	s := keys(m)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	if len(s) >= 2 {
		fmt.Println(s[0] * s[1])
	} else {
		fmt.Println(0)
	}

}

func keys(m map[int]int) []int {
	ks := []int{}
	for k, v := range m {
		if v >= 2 && v < 4 {
			ks = append(ks, k)
		}
		if 4 <= v {
			ks = append(ks, k)
			ks = append(ks, k)
		}
	}
	return ks
}
