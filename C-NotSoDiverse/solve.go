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
	k := nextInt()

	// ball := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	ball[i] = nextInt()
	// }

	m := map[int]int{}
	for i := 0; i < n; i++ {
		key := nextInt()
		if _, ok := m[key]; ok {
			m[key]++
		} else {
			m[key] = 1
		}
	}

	vslice := []int{}
	for _, v := range m {
		vslice = append(vslice, v)
	}
	sort.Sort(sort.IntSlice(vslice))

	// fmt.Println(m)
	// fmt.Println(vslice)

	count := 0
	for len(vslice) > k {
		count += vslice[0]
		vslice = vslice[1:]
	}

	fmt.Println(count)

}
