package main

import (
	"bufio"
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n)
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	m := map[int]int{}
	for i := 1; i <= n; i++ {
		m[i] = n / i
	}

	for i := n; i >= 1; i-- {
		if m[i] == 1 {
			b[i] = 1
			continue
		} else {
			b[i] = 0
			continue
		}

	}

}
