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
	m := map[int]int{}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	for i := 0; i < n; i++ {
		for a[i]%2 == 0 {
			a[i] /= 2
		}
		m[a[i]] = 1
	}

	fmt.Println(len(m))
}
