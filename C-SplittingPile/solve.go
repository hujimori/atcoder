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

	// if n == 2 {
	// 	fmt.Println(abs(a[0] - a[1]))
	// 	return
	// }

	left := make([]int, n+1)
	left[0] = 0
	for i := 0; i < n; i++ {
		left[i+1] = a[i] + left[i]
	}
	right := make([]int, n+1)
	right[0] = 0
	for i := 0; i < n; i++ {
		right[i+1] = a[n-1-i] + right[i]
	}

	min := math.Inf(1)

	for i := 1; i < n; i++ {
		x := left[i]
		y := right[n-i]

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
