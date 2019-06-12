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
	m := nextInt()

	city := make([]int, n)
	for i := 0; i < n; i++ {
		city[i] = 0
	}

	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = nextInt() - 1
		b[i] = nextInt() - 1
	}

	for i := 0; i < m; i++ {
		city[a[i]]++
		city[b[i]]++
	}

	for _, c := range city {
		fmt.Println(c)
	}
}
