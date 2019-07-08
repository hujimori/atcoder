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
	a := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		s += a[i]
	}

	x := make([]int, n)
	for i := 1; i < n; i += 2 {
		s -= 2 * a[i]
	}
	x[0] = s

	for i := 0; i < n-1; i++ {
		x[i+1] = 2*a[i] - x[i]
	}

	for _, v := range x {
		fmt.Printf("%d ", v)
	}

}
