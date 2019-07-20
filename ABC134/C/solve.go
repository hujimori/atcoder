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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		b[i] = a[i]
	}

	sort.Ints(b)
	// fmt.Println(b)
	// fmt.Println(a)
	for i := 0; i < n; i++ {
		if a[i] == b[n-1] {
			fmt.Println(b[n-2])
		} else {
			fmt.Println(b[n-1])
			// ans := upperBound(tmp, a[i])
		}
		// fmt.Println(tmp[ans-1])
	}

}

func upperBound(t []int, k int) int {
	lb := -1
	ub := len(t)
	for ub-lb > 1 {
		mid := (lb + ub) / 2

		if t[mid] <= k {
			lb = mid
		} else {
			ub = mid
		}
	}
	// if t[lb] <= k {
	// 	return ub
	// }
	return ub
}
