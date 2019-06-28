package main

import (
	"bufio"
	"fmt"
	"math"
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
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	if n =


	left := make([]int, n)
	left[0] = 0
	for i := 0; i < n-1; i++ {
		left[i+1] = a[i] + left[i]
	}

	fmt.Println(left)
	min := math.Inf(1)

	for i := 1; i < n-1; i++ {
		x := left[i]
		y := left[n-1] - left[i+1]
		fmt.Printf("x=%d, y=%d\n", x, y)

		rec := float64(abs(x - y))
		if min > rec {
			min = rec
		}
	}

	fmt.Println(int(min))

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
