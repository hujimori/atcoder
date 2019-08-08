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

var N, M int

func main() {
	sc.Split(bufio.ScanWords)

	N, M = nextInt(), nextInt()
	a := make([][]int, 3)

	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			a[j] = append(a[j], nextInt())
		}
	}

	res := 0
	for bit := uint(0); bit < (1 << 3); bit++ {
		var b []int
		for i := 0; i < N; i++ {
			tmp := 0
			for j := uint(0); j < 3; j++ {
				if bit&(1<<j) == 1<<j {
					tmp += a[j][i]
				} else {
					tmp -= a[j][i]
				}
			}
			b = append(b, tmp)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(b)))
		sum := 0
		for i := 0; i < M; i++ {
			sum += b[i]
		}
		res = max(res, sum)
	}

	fmt.Println(res)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
