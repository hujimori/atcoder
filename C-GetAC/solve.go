package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var n, q int
	fmt.Scan(&n, &q)

	// n, q := nextInt(), nextInt()
	// n++
	var s string
	fmt.Scan(&s)
	slice := strings.Split(s, "")

	l := make([]int, q)
	r := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&l[i], &r[i])
	}

	ans := make([]int, n)
	ans[0] = 0

	for i := 0; i < n-1; i++ {
		if (slice[i] + slice[i+1]) == "AC" {
			ans[i+1] = ans[i] + 1
		} else {
			ans[i+1] = ans[i]
		}
	}

	for i := 0; i < q; i++ {
		right := ans[r[i]-1]
		left := ans[l[i]-1]
		fmt.Println(right - left)
	}
}
