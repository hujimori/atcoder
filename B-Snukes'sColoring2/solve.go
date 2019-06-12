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

	w := nextInt()
	h := nextInt()
	n := nextInt()

	x := make([]int, n)
	y := make([]int, n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		x[i] = nextInt()
		y[i] = nextInt()
		a[i] = nextInt()
	}

	wh := make([][]int, h)
	for i := 0; i < h; i++ {
		wh[i] = make([]int, w)
		for j := 0; j < w; j++ {

			wh[i][j] = 0
		}

	}

	for i := 0; i < n; i++ {
		if a[i] == 1 {
			for j := 0; j < h; j++ {
				for k := 0; k < x[i]; k++ {
					wh[j][k] = 1
				}
			}
		}

		if a[i] == 2 {
			for j := 0; j < h; j++ {
				for k := x[i]; k < w; k++ {
					wh[j][k] = 1
				}
			}
		}

		if a[i] == 3 {
			for j := 0; j < y[i]; j++ {
				for k := 0; k < w; k++ {
					wh[j][k] = 1
				}
			}
		}

		if a[i] == 4 {
			for j := y[i]; j < h; j++ {
				for k := 0; k < w; k++ {
					wh[j][k] = 1
				}
			}
		}
	}

	// for i := 0; i < h; i++ {
	// 	fmt.Println(wh[i])
	// }

	area := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if wh[i][j] == 0 {
				area++
			}
		}
	}

	fmt.Println(area)

}
