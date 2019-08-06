package main

import (
	"bufio"
	"fmt"
	"os"
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

	ans := -100000000
	for t := 0; t < len(a); t++ {
		takahashi := -100000000
		aoki := -100000000

		for ao := 0; ao < len(a); ao++ {
			if t == ao {
				continue
			}

			l := min(t, ao)
			r := max(t, ao)
			len := r - l + 1
			ts := 0
			as := 0
			for i := 0; i < len; i++ {
				if i%2 == 0 {
					ts += a[i+l]
				} else {
					as += a[i+l]
				}
			}

			if aoki < as {
				aoki = as
				takahashi = ts
			}
		}

		ans = max(ans, takahashi)
	}

	fmt.Println(ans)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
