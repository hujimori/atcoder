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
	n, x := nextInt(), nextInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	sum := make([]int, n)
	for i := 0; i < n-1; i++ {
		sum[i] = a[i] + a[i+1]
	}

	// cost := 0
	// for i := 0; i < len(sum); i++ {
	// 	if sum[i] > x {
	// 		c := sum[i] - x
	// 		sum[i] -= c
	// 		sum[i+1] -= c
	// 		cost += c
	// 	}
	// 	// fmt.Println(sum)
	// }

	// fmt.Println(cost)

	count := 0
	for i := 0; i < n-1; i++ {
		if a[i]+a[i+1] > x {
			c := a[i] + a[i+1] - x
			count += c
			if a[i+1] > c {
				a[i+1] -= c
			} else {
				c -= a[i+1]
				a[i+1] = 0
				a[i] -= c
			}

			// if a[i] > a[i+1] {
			// 	a[i] -= c
			// 	count += c
			// } else if a[i] < a[i+1] {
			// 	a[i+1] -= c
			// 	count += c
			// } else {
			// 	a[i+1] -= c
			// 	count += c
			// }

		}
	}

	// fmt.Println(a)

	fmt.Println(count)
}
