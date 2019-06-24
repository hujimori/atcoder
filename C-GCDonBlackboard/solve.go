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
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	left := make([]int, n+1)
	right := make([]int, n+1)

	for i := 0; i < n; i++ {
		left[i+1] = gcd(left[i], a[i])
	}

	for i := n - 1; i >= 0; i-- {
		right[i] = gcd(right[i+1], a[i])
	}

	res := 0.0
	for i := 0; i < n; i++ {
		l := left[i]

		r := right[i+1]

		chmax(&res, float64(gcd(l, r)))
	}
	fmt.Println(int(res))

	// slice := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	slice[i] = nextInt()
	// }

	// max := -1

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == j {
	// 			continue
	// 		}

	// 		res := gcd(slice[i], slice[j])

	// 		if max < res {
	// 			max = res
	// 		}
	// 	}
	// }
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a

}

func chmin(a *float64, b float64) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax(a *float64, b float64) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
