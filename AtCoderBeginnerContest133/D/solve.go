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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	m := make([]int, n+1)
	m[0] = 0
	// m[1] = 2 * a[0]

	for i := 1; i < n+1; i++ {
		m[i] = 2*a[i-1] - m[i-1]
	}

	fmt.Println(m)

}
